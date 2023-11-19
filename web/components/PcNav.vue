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
        <li class="mr-2">
          <button class="shadow-center rounded-full w-10 h-10">
            <KLink to="/">
              <Icon name="clarity:house-line" size="1.5rem"></Icon>
            </KLink>
          </button>
        </li>
        <li class="mr-2">
          <button class="shadow-center rounded-full w-10 h-10">
            <Icon name="bi:journal-richtext" size="1.5rem"></Icon>
          </button>
        </li>
        <li class="mr-2">
          <button class="shadow-center rounded-full w-10 h-10">
            <Icon name="ic:round-ondemand-video" size="1.5rem"></Icon>
          </button>
        </li>
        <li class="mr-2">
          <button class="shadow-center rounded-full w-10 h-10">
            <Icon name="bi:journal-text" size="1.5rem"></Icon>
          </button>
        </li>
      </ul>
    </div>
    <div class="md:col-span-1 flex items-center justify-end">
      <KButton @click="toCreate">发表</KButton>
      <!-- <SwitcherTheme /> -->

      <div class="ml-4 w-12 h-12">
        <KLink v-if="isLogin" :to="getPath()">
          <Avatar :url="userInfo.avatar" class="w-full h-full" />
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
const store = useLoginStore();
const userStore = useUserStore();
const userInfo = computed(() => userStore.getUserInfo);
const isLogin = computed(() => userStore.getIsLogin);
const { to } = useToRoute();
const setLoginStatus = () => {
  store.setLoginStatus();
};

const toCreate = () => {
  if (isLogin.value) {
    to(`/topic/create`);
    return;
  }
  addMessage("请先登录", "warning");
};
function getPath() {
  console.log(userInfo);
  return `/user/${userInfo.value.id}`;
}
userStore.fetchUserInfo();
</script>
