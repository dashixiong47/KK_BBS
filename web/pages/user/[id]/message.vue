<template>
  <div class="w-full mb-5 flex justify-end">
    <KButton @click="readAll">一键已读</KButton>
  </div>
  <ul>
    <li
      v-for="item in messageList"
      class="shadow-center rounded-xl p-5 mb-5 flex justify-between items-end"
    >
      <div>
        <h3 class="text-lg mb-2 relative flex items-center">
          <Avatar :url="item.relatedUser.avatar" class="w-10 h-10" />
          <div class="flex items-center">
            <span class="ml-2">{{ item.relatedUser.nickname }}</span>
            <span class="ml-2 text-sm">
              {{ getRelativeTime(item.createdAt) }}
            </span>
            <span class="ml-2">{{ typeTitle[item.type] }}</span>
          </div>
          <Icon
            v-if="!item.isView"
            class="text-red-600 -mt-5"
            name="material-symbols:fiber-new-outline"
          ></Icon>
        </h3>
        <h4 class="text-base mb-2">《{{ item.topicTitle }}》</h4>
        <p class="details" v-if="[2, 3].includes(item.type)">
          " <span>{{ item.content }} </span>"
        </p>
      </div>
      <KLink v-if="item.detail" :to="item.detail" class="flex-shrink-0 ml-5">
        <KButton @click.stop="change(item.id)">查看详情</KButton>
      </KLink>
    </li>
  </ul>
  <Pagination v-model="pages" @query="init" />
</template>

<script setup>
import { readMessage, readMessageAll } from "@/api/index";
import { useAppConfigStore } from "~/stores/init";
import { useMessageStore } from "~/stores/message.js";
const { getRelativeTime } = useTime();
const messageStore = useMessageStore();
let appConfigStore = useAppConfigStore();
let appConfig = computed(() => appConfigStore.getConfig);
useHead({
  title: "消息 - " + appConfig.value.appName,
});
let pages = ref({
  total: 0,
  page: 1,
  pageSize: 10,
});
const messageList = computed(() => messageStore.getMessageData);
const messageTotal = computed(() => messageStore.getMessageTotal);
watch(
  () => messageTotal.value,
  (newV) => {
    pages.value.total = newV;
  },
  {
    immediate: true,
  }
);

const typeTitle = {
  1: "系统通知",
  2: "评论了您的帖子",
  3: "回复了您的评论",
  4: "点赞了您的帖子",
  6: "收藏了您的帖子",
};
const change = async (id) => {
  try {
    await readMessage(id);
  } catch (error) {}
};
const readAll = async () => {
  if (!messageStore.getUnreadMessage) return;
  try {
    await readMessageAll();
    init();
    messageStore.setUnreadMessage(0);
  } catch (error) {}
};
const init = () => {
  messageStore.fetchMessageList(pages);
};
init();
</script>
