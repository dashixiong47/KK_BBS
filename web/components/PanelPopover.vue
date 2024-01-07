<template>
  <div class="p-5 shadow-center flex flex-col items-center rounded-xl">
    <div class="text-xl">{{ userInfo.nickname }}</div>
    <!-- {{ userInfo }} -->
    <div class="w-full flex flex-col items-center my-2">
      <span class="text-sm">
        金币:
        <span class="text-red-500">
          {{ userInfo.coins }}
          <Icon name="healthicons:coins-outline"></Icon>
        </span>
      </span>
      <div class="w-full flex justify-between mt-2">
        <KLink :to="getPath" class="flex flex-col items-center">
          <span class="font-bold">{{ userInfo.topicTotal }}</span>
          <span class="text-xs">动态</span>
        </KLink>
        <KLink :to="getPath + '/follow'" class="flex flex-col items-center">
          <span class="font-bold">{{ userInfo.follow }}</span>
          <span class="text-xs">关注</span>
        </KLink>
        <KLink :to="getPath + '/fans'" class="flex flex-col items-center">
          <span class="font-bold">{{ userInfo.fans }}</span>
          <span class="text-xs">粉丝</span>
        </KLink>
      </div>
    </div>
    <div class="w-full">
      <div
        v-for="item in list"
        @click="item.event"
        class="w-full p-2 cursor-pointer rounded-xl flex items-center justify-between hover:text-white hover:bg-blue-400"
      >
        <div>
          <Icon size="1.2rem" :name="item.iconName" class="mr-1"></Icon>
          <span>{{ item.name }}</span>
        </div>
        <Icon
          v-if="item.rightIconName"
          size="1.2rem"
          :name="item.rightIconName"
          class="mr-1"
        ></Icon>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useLoginStore, useUserStore } from "~/stores/main.js";
const { setCookie } = useCookies();
// 使用状态管理库
const loginStore = useLoginStore();
const userStore = useUserStore();
const props = defineProps({
  userInfo: {
    type: Object,
    default: () => ({}),
  },
});
const list = [
  {
    name: "个人中心",
    iconName: "ep:user",
    rightIconName: "ep:arrow-right",
    event: () => {},
  },
  {
    name: "个人中心",
    iconName: "ep:user",
    rightIconName: "ep:arrow-right",
    event: () => {},
  },
  {
    name: "退出登录",
    iconName: "ep:switch-button",
    rightIconName: "",
    event: () => {
      setCookie("token", "");
      userStore.setIslogin(false);
    },
  },
];
const getPath = computed(() => {
  return `/user/${props.userInfo.id}`;
});
</script>
<style scoped></style>
