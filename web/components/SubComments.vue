<template>
  <ul class="mt-2">
    <li v-for="item in replyList" class="flex mb-2 pb-2">
      <Avatar class="w-6 h-6" :url="item.user?.avatar" />
      <div class="ml-2 w-full">
        <div class="text-sm">
          <span class="mb-2">
            <KName :name="item.user.nickname" />
            <span class="flex-shrink-0 mx-2">
              {{ $t("comments-reply") }}
            </span>
            <KName
              :name="`@${item.replyUser.nickname}`"
              class="text-blue-400"
            />
            :
          </span>
          {{ item.content }}
        </div>
        <div class="text-xs text-light-5">
          <div class="flex">
            <span class="mr-2 cursor-pointer">{{
              getRelativeTime(item.createdAt)
            }}</span>
            <span
              class="mr-2 cursor-pointer flex items-center"
              :class="{
                'text-[--color-primary]': item.likeState,
                'dark:text-[--dark-color-primary]': item.likeState,
              }"
              @click="commentLikeChange(item, parentId, item.id)"
            >
              <Icon name="icon-park-solid:thumbs-up" class="mr-1" />{{
                item.like
              }}
            </span>
            <span
              class="cursor-pointer"
              @click="(e) => subReply(e, item, parentId)"
              >{{ $t("comments-reply") }}
            </span>
          </div>
          <div class="reply"></div>
        </div>
      </div>
    </li>
    <p
      @click="loadMore"
      v-if="total > 3 && replyList.length < total"
      class="text-xs secondary-text cursor-pointer"
    >
      剩余{{ total - replyList.length }}条回复, 点击查看
    </p>
  </ul>
</template>

<script setup>
import { subCommentList } from "@/api/index";
const { addMessage } = useMessage();
const { getRelativeTime } = useTime();
let { reply, subReply, commentLikeChange, parentId, total, topicId } =
  defineProps({
    reply: {
      type: Array,
      default: () => [],
    },
    subReply: {
      type: Function,
      default: () => {},
    },
    commentLikeChange: {
      type: Function,
      default: () => {},
    },
    parentId: {
      type: String,
      default: "",
    },
    topicId: {
      type: String,
      default: "",
    },
    total: {
      type: Number,
      default: 0,
    },
  });
let page = {
  page: 1,
  pageSize: 10,
};
let subList = ref([]);
let replyList = computed(() => {
  if (!reply) return [];
  if (subList.value.length) {
    return subList.value;
  } else {
    return reply;
  }
});
async function loadMore() {
  try {
    let data = await subCommentList(parentId, { topicId, ...page });
    subList.value = subList.value.concat(data);
    page.page++;
  } catch (error) {
    addMessage({
      type: "error",
      content: error,
    });
  }
}
</script>
