<template>
  <div class="shadow-center rounded-2xl relative overflow-hidden flex flex-col">
    <!-- 背景图片和头像 -->
    <div class="pt-[33%] sm:md:pt-[20%] md:pt-[16%] h-2 flex-shrink-0 relative">
      <img
        class="object-cover absolute inset-0 w-full h-full"
        :src="userInfo.background || '/images/bg.png'"
        alt=""
        srcset=""
      />
      <div
        class="overflow-hidden rounded-full absolute w-14 h-14 sm:w-20 sm:h-20 md:w-24 md:h-24 z-10 -bottom-4 left-1/2 -translate-x-1/2"
      >
        <Avatar
          :url="userInfo.avatar"
          class="avatar shadow-md bg-blue-400 w-full h-full"
        />
        <div
          class="edit text-xs w-full h-full absolute bottom-0 left-1/2 -translate-x-1/2 z-0 flex items-center justify-center text-white cursor-pointer"
        >
          <Cropping @uploadSuccess="uploadSuccess">
            <Icon
              size="1rem"
              name="material-symbols-light:edit-square-outline-rounded"
            ></Icon>
            更换头像
          </Cropping>
        </div>
      </div>
    </div>

    <!-- 用户信息区域 -->
    <div class="pt-5 pb-5 flex flex-col items-center">
      <!-- 昵称 -->
      <h2 class="nickname flex items-center font-bold mb-2">
        <span v-if="nicknameState">{{ userInfo.nickname }}</span>
        <KInput
          class="p-2 rounded-lg"
          v-else
          v-model="userInfo.nickname"
          @keyup.enter="
            updateAndHideInput({ nickname: userInfo.nickname }, 'nickname')
          "
        />
        <span v-if="nicknameState" class="nicknameEdit" @click="showNickname">
          <Icon name="material-symbols-light:edit-square-outline-rounded" />
        </span>
      </h2>
      <!-- 介绍 -->
      <p class="introduction flex items-center text-xs">
        <span v-if="introductionState">{{ userInfo.introduction }}</span>
        <KInput
          class="p-2 rounded-lg"
          v-else
          v-model="userInfo.introduction"
          @keyup.enter="
            updateAndHideInput(
              { introduction: userInfo.introduction },
              'introduction'
            )
          "
        />
        <span
          v-show="introductionState"
          class="introductionEdit"
          @click="showIntroduction"
        >
          <Icon name="material-symbols-light:edit-square-outline-rounded" />
        </span>
      </p>
    </div>
    <!-- Tab栏 -->
    <div>
      <ul class="flex items-center border-t px-5">
        <li
          v-for="(item, index) in tabs"
          class="h-12 flex items-center mx-5 cursor-pointer"
          :class="[activeTab(item.path) ? 'border-b-2 border-blue-500' : '']"
        >
          <KLink :to="item.path" class="flex items-center">
            <span class="text-sm">{{ item.name }}</span>
            <!-- <span class="ml-1 text-xs text-gray-400">{{
              formatNumber(userInfo[item.name.toLowerCase()])
            }}</span> -->
          </KLink>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { updateUserInfo } from "~/api";
import { useGetUserInfo } from "~/api/server";
import { useUserStore } from "~/stores/main";
const userStore = useUserStore();
const { addMessage } = useMessage();
const route = useRoute();
let tabs = [
  {
    name: "动态",
    path: "/user/" + route.params.id,
    active: "",
  },
  {
    name: "关注",
    path: "/user/" + route.params.id + "/follow",
    active: "follow",
  },
  {
    name: "粉丝",
    path: "/user/" + route.params.id + "/fans",
    active: "fans",
  },
];
let userInfo = ref({});
let nicknameState = ref(true);
let introductionState = ref(true);

let activeTab = (path) => {
  return path === route.path;
};
const showNickname = () => {
  nicknameState.value = false;
};

const showIntroduction = () => {
  introductionState.value = false;
};

function uploadSuccess(item) {
  update({ avatar: item.url });
}

async function update(data = null) {
  if (!data) return;
  try {
    await updateUserInfo(data);
    addMessage("更新成功", "success");
    userStore.setUserInfo({
      ...userInfo.value,
      ...data,
    });
    init();
  } catch (error) {
    // 处理错误
    console.error(error);
  }
}

async function updateAndHideInput(data, type) {
  await update(data);
  if (type === "nickname") {
    nicknameState.value = true;
  } else if (type === "introduction") {
    introductionState.value = true;
  }
}

async function init() {
  try {
    const { data } = await useGetUserInfo(route.params.id);
    userInfo.value = data;
  } catch (error) {
    // 处理错误
    console.error(error);
  }
}

init();
</script>

<style lang="postcss">
.avatar:hover + .edit {
  display: flex;
}
.edit {
  display: none;
  background: rgba(0, 0, 0, 0.5);
  border-radius: 50%;
  transition: all 0.3s;
}
.edit:hover {
  display: flex;
}
.introduction:hover .introductionEdit,
.nickname:hover .nicknameEdit {
  display: inline-block;
}
.introductionEdit,
.nicknameEdit {
  cursor: pointer;
  display: none;
}
</style>
