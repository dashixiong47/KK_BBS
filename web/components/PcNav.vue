<template>
  <nav
    class="glass shadow-md border-b h-16 box-content flex-shrink-0 content-center px-5 grid grid-cols-2 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4"
  >
    <div class="md:col-span-1 primary-text">
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
        class="w-4/6 h-full flex items-center px-5 text-[--color-primary] dark:text-[--dark-color-primary]"
      >
        <li class="mr-2">
          <button class="glass border rounded-full w-10 h-10">
            <KLink to="/">
              <Icon name="icon-zhuye" size="text-2xl mx-1"></Icon>
            </KLink>
          </button>
        </li>
        <li class="mr-2">
          <button class="glass border rounded-full w-10 h-10">
            <Icon name="icon-neirong" size="text-2xl mx-1"></Icon>
          </button>
        </li>
        <li class="mr-2">
          <button class="glass border rounded-full w-10 h-10">
            <Icon name="icon-video" size="text-2xl mx-1"></Icon>
          </button>
        </li>
        <li class="mr-2">
          <button class="glass border rounded-full w-10 h-10">
            <Icon name="icon-wendang" size="text-2xl mx-1"></Icon>
          </button>
        </li>
      </ul>
    </div>
    <div class="md:col-span-1 flex items-center justify-end">
      <KLink :to="`/topic/create`">发表</KLink>
      <Icon @click="test" name="icon-pinglun" size="text-2xl"></Icon>
      <Icon  name="icon-xiaoxi" size="text-2xl mx-2"></Icon>
      <Icon name="icon-shezhi" size="text-2xl mx-2"></Icon>
      <Switcher />
      <KLink class="ml-4  w-11 h-11" :to="`/user/1`">
        <Avatar v-if="userInfo.token" :url="userInfo.avatar" class="w-full h-full" />
        <button v-else @click="setLoginStatus">登录/注册</button>
      </KLink>
      <!-- <div class="absolute">1212</div> -->
    </div>
  </nav>
  <teleport to="body">
    <Login />
  </teleport>
</template>

<script setup>
import { useLoginStore, useUserStore } from "~/stores/main.js";
const store = useLoginStore();
const userStore = useUserStore();

const userInfo = computed(() => userStore.getUserInfo);
const setLoginStatus = () => {
  store.setLoginStatus();
};

const { notice } = useNotice();

const test = () => {
  notice({
    title: "标题",
    content:
      "一、放下大概就是这样，即使我们没在一起，我也会好好的，谢谢时间惊艳了那段有你的记忆，也谢谢现在更努力变好的自己。",
    autoClose: true,
  });
};
onMounted(async () => {
  try {
    let userInfoStr = localStorage.getItem("userInfo");
    if (userInfoStr) {
      userStore.setUserInfo(JSON.parse(userInfoStr));
    }
    // let res = await getUserInfo();
    // console.log(res);
  } catch (error) {}
});
</script>
