<template>
  <nav
    class="glass z-10 shadow-md border-b h-16 box-content flex-shrink-0 content-center px-5 grid grid-cols-2 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4"
  >
    <div class="md:col-span-1 font-main-color">
      <KLink to="/" class="flex items-center h-full">
        <img
          class="h-8 mr-5"
          src="https://hxsj.in/_nuxt/img/logo.5c704a9.png"
          alt=""
          srcset=""
        />
        KK_BBS
      </KLink>
    </div>
    <div class="hidden lg:flex items-center col-span-2">
      <Search />
      <ul
        class="w-4/6 h-full flex items-center px-5 text-[--illuminate-color] dark:text-[--dark-illuminate-color]"
      >
        <li class="mr-2" v-for="item in type">
          <button class="shadow-center rounded-full w-10 h-10">
            <KLink :to="item.path">
              <Icon :name="item.icon" size="1.5rem"></Icon>
            </KLink>
          </button>
        </li>
      </ul>
    </div>
    <div class="md:col-span-1 flex items-center justify-end">
      <KButton @click="toCreate">发表</KButton>
      <!-- <SwitcherTheme /> -->

      <div class="ml-4 flex-shrink-0">
        <KLink v-if="isLogin" :to="getPath()">
          <Avatar :url="userInfo.avatar" class="w-12 h-12" />
        </KLink>
        <KButton v-else @click="setLoginStatus">登录</KButton>
      </div>
      <!-- <div class="absolute">1212</div> -->
    </div>
  </nav>
  <Login />
</template>

<script setup>
import { useLoginStore, useUserStore } from "~/stores/main.js";
const { addMessage } = useMessage();
const store = useLoginStore();
const userStore = useUserStore();
const userInfo = computed(() => userStore.getUserInfo);
const isLogin = computed(() => userStore.getIsLogin);
const { to } = useToRoute();
const setLoginStatus = () => {
  store.setLoginStatus();
};
const type = [
  {
    name: "首页",
    icon: "clarity:house-line",
    path: "/",
  },
  {
    name: "默认",
    icon: "icon-park-solid:topic",
    path: "/topic",
  },
  {
    name: "图片",
    icon: "lets-icons:img-box-fill",
    path: "/img",
  },
  {
    name: "视频",
    icon: "ic:round-ondemand-video",
    path: "/video",
  },
  {
    name: "小说",
    icon: "mdi:text-box-multiple",
    path: "/text",
  },
];
const toCreate = () => {
  if (isLogin.value) {
    to(`/create`);
    return;
  }
  addMessage("请先登录", "warning");
  setLoginStatus();
};
function getPath() {
  return `/user/${userInfo.value.id}`;
}
userStore.fetchUserInfo();
</script>
