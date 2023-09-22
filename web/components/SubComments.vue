<template>
  <ul class="mt-2">
    <li v-for="item in reply" class="flex mb-2 pb-2">
      <Avatar class="w-6 h-6" :url="item.user?.avatar" />
      <div class="ml-2 w-full">
        <div class="text-sm">
          <span class="mb-2">
            <KName :name="item.user.nickname" />
            <span class="flex-shrink-0 mx-2">
              {{ $t("subComments-reply") }}
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
          <span class="mr-2 cursor-pointer">{{
            getRelativeTime(item.createdAt)
          }}</span>
          <span class="mr-2 cursor-pointer">
            <Icon name="icon-park-solid:thumbs-up" class="mr-1" />11
          </span>
          <span
            class="cursor-pointer"
            @click="(e) => subReply(e, item, parentId)"
            >{{ $t("subComments-reply") }}</span
          >
          <div class="reply"></div>
        </div>
      </div>
    </li>
    <p v-if="total>3" class="text-xs secondary-text cursor-pointer">共{{ total }}条回复, 点击查看</p>
  </ul>
</template>

<script setup>
const { getRelativeTime } = useTime();
let { reply, subReply, parentId, total } = defineProps({
  reply: {
    type: Array,
    default: () => [],
  },
  subReply: {
    type: Function,
    default: () => {},
  },
  parentId: {
    type: String,
    default: "",
  },
  total: {
    type: Number,
    default: 0,
  },
});
</script>
