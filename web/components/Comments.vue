<template>
  <p class="mb-2 w-full">评论</p>
  <ReplyInput v-model="replyData" @success="replySuccess" />
  <ReplyInput
    class="mt-2"
    v-show="showSubReply"
    ref="subReplyIpt"
    v-model="subReplyData"
    @success="replySuccess"
  />
  <div class="flex items-center mb-5">
    <div class="flex items-center">
      <span class="text-base font-extrabold">{{ $t("comments-title") }}</span>
      <span class="ml-2 text-sm font-normal text-light-5">{{ total }}</span>
    </div>
    <div class="ml-5 text-sm text-light-5">
      <button
        class="border px-2 rounded-l-md"
        :class="{
          'bg-light-3 dark:bg-light-4': !selected,
          'text-blue-400': !selected,
        }"
        @click="change(0)"
      >
        {{ $t("comments-hot") }}
      </button>
      <button
        class="border px-2 rounded-r-md"
        :class="{
          'bg-light-3 dark:bg-light-4': selected,
          'text-blue-400': selected,
        }"
        @click="change(1)"
      >
        {{ $t("comments-new") }}
      </button>
    </div>
  </div>
  <ul>
    <li
      v-for="(item, index) in comments"
      class="flex mb-5 pb-5"
      :class="{
        'border-b': index !== comments.length - 1,
      }"
    >
      <Avatar :url="item.user?.avatar" />
      <div class="ml-5 w-full">
        <p class="mb-2"><KName :name="item.user?.nickname" /></p>
        <div class="text-sm">
          {{ item.content }}
        </div>
        <div class="text-xs text-light-5 flex items-center">
          <span class="mr-2 cursor-pointer">{{
            getRelativeTime(item.createdAt)
          }}</span>
          <span class="mr-2 cursor-pointer">
            <Icon name="icon-park-solid:thumbs-up" class="mr-1" />11
          </span>
          <span class="cursor-pointer" @click="(e) => subReply(e, item,item.id)">{{
            $t("comments-reply")
          }}</span>
        </div>
        <div class="reply"></div>
        <div>
          <SubComments :total="item.total" :parentId="item.id" :reply="item.reply" :subReply="subReply" />
        </div>
      </div>
    </li>
  </ul>
  <div>
    <KButton
      class="w-full flex justify-center"
      :disabled="getDisabled"
      :throttle="loadMore"
      :class="{ 'cursor-not-allowed': getDisabled }"
      >{{ getDisabled ? "没有更多了" : "加载更多" }}</KButton
    >
  </div>
</template>

<script setup>
import { useGetComments } from "~/api/server";
const { notice } = useNotice();
const { getRelativeTime } = useTime();
let { topicId } = defineProps({
  topicId: {
    type: String,
    required: true,
  },
});
let showSubReply = ref(false);
let comments = ref([]);
let total = ref(0);
let replyData = ref({
  topicId: topicId,
  content: null,
});
let subReplyData = ref({
  topicId: topicId,
  parentId: null,
  replyToUserId: null,
  content: null,
});
let page = {
  page: 1,
  pageSize: 10,
};
let selected = ref(0);
let subReplyIpt = ref(null);
const getDisabled = computed(() => {
  return comments.value.length === total.value;
});
const change = (num) => {
  selected.value = num;
};
const replySuccess = () => {
  showSubReply.value = false;
  getComments();
};
const subReply = (e, item,parentId) => {
  subReplyData.value.parentId = parentId;
  subReplyData.value.replyToUserId = item.user.id;
  let replyNode = e.target.parentNode.parentNode.querySelector(".reply");
  let subReplyNode = subReplyIpt.value.$el;
  showSubReply.value = true;
  replyNode.appendChild(subReplyNode);
};
const loadMore = () => {
  if (page.page * page.pageSize >= total.value) {
    notice({
      title: "提示",
      content: "没有更多了",
    });
    return;
  }
  page.page++;
  getComments();
};
async function getComments() {
  try {
    let data = await useGetComments(topicId, page);
    
    // 计算当前页面在 comments.value 数组中的起始和结束索引
    const startIndex = (page.page - 1) * page.pageSize;
    const endIndex = startIndex + page.pageSize;
    
    // 如果 comments.value 有数据，则替换；否则，连接。
    if (comments.value.length > 0) {
      // 删除旧的数据范围，并插入新的数据
      comments.value.splice(startIndex, endIndex - startIndex, ...data.comments);
    } else {
      // 如果 comments.value 是空的，直接连接新数据
      comments.value = comments.value.concat(data.comments);
    }
    
    total.value = data.total;
  } catch (error) {
    console.log(error);
  }
}

function init() {
  getComments();
}
init();
</script>
