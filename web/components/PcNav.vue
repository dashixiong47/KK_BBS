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
              <Icon name="clarity:house-line" size="1.5rem"></Icon>
            </KLink>
          </button>
        </li>
        <li class="mr-2">
          <button class="glass border rounded-full w-10 h-10">
            <Icon name="bi:journal-richtext" size="1.5rem"></Icon>
          </button>
        </li>
        <li class="mr-2">
          <button class="glass border rounded-full w-10 h-10">
            <Icon name="ic:round-ondemand-video" size="1.5rem"></Icon>
          </button>
        </li>
        <li class="mr-2">
          <button class="glass border rounded-full w-10 h-10">
            <Icon name="bi:journal-text" size="1.5rem"></Icon>
          </button>
        </li>
      </ul>
    </div>
    <div class="md:col-span-1 flex items-center justify-end">
      <KLink :to="`/topic/create`">发表</KLink>
      <Icon
        @click="test"
        name="icon-park-solid:all-application"
        size="1rem"
        class="mr-2"
      />
      <Icon name="icon-park-solid:all-application" size="1rem" class="mr-2" />
      <Icon name="icon-park-solid:all-application" size="1rem" class="mr-2" />
      
      <SwitcherTheme />
    
      <div class="ml-4 w-11 h-11">
        <KLink v-if="Object.keys(userInfo).length" :to="`/user/1`">
          <Avatar :url="userInfo.avatar" class="w-full h-full" />
        </KLink>
        <button v-else @click="setLoginStatus">登录/注册</button>
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
const loginStatus= computed(()=>store.getLoginStatus)
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
userStore.fetchUserInfo();
onMounted(async () => {
  // try {
  //   let userInfoStr = localStorage.getItem("userInfo");
  //   if (userInfoStr) {
  //     userStore.setUserInfo(JSON.parse(userInfoStr));
  //   }
  //   // let res = await getUserInfo();
  //   // console.log(res);
  // } catch (error) {}
});
</script>
