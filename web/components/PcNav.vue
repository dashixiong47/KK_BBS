<template>
  <nav
    class="z-10 shadow-md border-b h-16 box-content flex-shrink-0 content-center px-5 grid grid-cols-2 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4"
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
          <KLink :to="item.path">
            <button
              class="shadow-center rounded-full w-10 h-10"
              :class="[
                item.path === route.path
                  ? 'bg-[--illuminate-color] dark:bg-[--dark-illuminate-color] text-white'
                  : 'bg-white dark:bg-[--dark-bg] text-[--illuminate-color] dark:text-[--dark-illuminate-color]',
              ]"
            >
              <Icon :name="item.icon" size="1.5rem"></Icon>
            </button>
          </KLink>
        </li>
      </ul>
    </div>
    <div class="md:col-span-1 flex items-center justify-end">
      <div class="flex-shrink-0">
        <div v-if="isLogin" class="relative options">
          <Tippy placement="bottom">
            <Avatar
              :showUserInfo="false"
              :url="userInfo.avatar"
              class="w-12 h-12 relative z-[10000] avatar"
            />
            <template #content>
              <PanelPopover :userInfo="userInfo" class="w-64 -mt-5" />
            </template>
          </Tippy>
        </div>
        <KButton v-else @click="setLoginStatus" class="mr-5">登录</KButton>
      </div>
      <KLink v-if="isLogin" :to="getPath + '/message'" class="relative">
        <Icon
          size="1.8rem"
          :name="
            unreadMessage
              ? 'ic:twotone-notifications-active'
              : 'ic:twotone-notifications'
          "
          class="mx-4"
          :class="{ message: unreadMessage }"
        ></Icon>
      </KLink>
      <KButton @click="toCreate">投稿</KButton>
      <!-- <SwitcherTheme /> -->

      <!-- <div class="absolute">1212</div> -->
    </div>
  </nav>
  <Login />
</template>

<script setup>
import { useLoginStore, useUserStore } from "~/stores/main.js";
import { useMessageStore } from "~/stores/message.js";
import "tippy.js/animations/scale.css";
const messageStore = useMessageStore();
const { addMessage } = useMessage();
const loginStore = useLoginStore();
const userStore = useUserStore();
const userInfo = computed(() => userStore.getUserInfo);
const isLogin = computed(() => userStore.getIsLogin);
const { to } = useToRoute();
const setLoginStatus = () => {
  loginStore.setLoginStatus();
};
const route = useRoute();
const unreadMessage = computed(() => messageStore.getUnreadMessage);
watch(
  () => route.path,
  (val) => {
    console.log(val);
  }
);
const type = [
  {
    name: "首页",
    icon: "ep:house",
    path: "/",
  },
  {
    name: "默认",
    icon: "ep:chat-line-round",
    path: "/topic",
  },
  {
    name: "图片",
    icon: "ep:picture",
    path: "/image",
  },
  {
    name: "视频",
    icon: "octicon:video-24",
    path: "/video",
  },
  {
    name: "小说",
    icon: "ep:document",
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
const getPath = computed(() => {
  return `/user/${userInfo.value.id}`;
});
userStore.fetchUserInfo();
messageStore.fetchUnreadMessage();
</script>

<style scoped>
.message {
  animation: shaking 0.3s infinite ease-in-out;
}

@keyframes shaking {
  0%,
  100% {
    transform: rotate(-5deg);
  }
  50% {
    transform: rotate(5deg);
  }
}
.options:hover .avatar {
  position: relative;
  top: 0;
  transform: scale(1.2);
  transition: all 0.3s ease-in-out;
}
/* .popover {
  transition: opacity 0.3s ease-in-out;
  width: 250px;
  opacity: 0;
  pointer-events: none;
} */
/* .options:hover .popover {
  opacity: 1;
  pointer-events: all;
} */
</style>
