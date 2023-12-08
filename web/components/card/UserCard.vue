<template>
  <div>
    <div class="shadow-center rounded-2xl p-5">
      <p class="border-b pb-2 mb-2">个人成就</p>
      <div class="flex flex-row justify-between text-sm">
        <div class="flex flex-col items-center">
          <span>{{ userStore.getUserInfo.topicTotal }}</span>
          <span>帖子</span>
        </div>
        <div class="flex flex-col items-center">
          <span>{{ userStore.getUserInfo.coins }}</span>
          <span>金币</span>
        </div>
        <div class="flex flex-col items-center">
          <span>{{ userStore.getUserInfo.follow }}</span>
          <span>关注</span>
        </div>
        <div class="flex flex-col items-center">
          <span>{{ userStore.getUserInfo.fans }}</span>
          <span>粉丝</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useUserStore } from "~/stores/main";

const userStore = useUserStore();
console.log(userStore.getUserInfo);
</script>
