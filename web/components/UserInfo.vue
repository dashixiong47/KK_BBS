<template>
  <div class="shadow-center rounded-2xl relative overflow-hidden flex flex-col">
    <div class="pt-[33%] sm:md:pt-[20%] md:pt-[16%] h-2 flex-shrink-0 relative">
      <img
        class="object-cover absolute inset-0 w-full h-full"
        :src="userInfo.background || '/images/bg.png'"
        alt=""
        srcset=""
      />
      <Avatar
        :url="userInfo.avatar"
        class="shadow-md bg-blue-400 absolute w-14 h-14 sm:w-20 sm:h-20 md:w-24 md:h-24 z-10 -bottom-4 left-1/2 -translate-x-1/2"
      />
    </div>

    <div class="pt-5 pb-5 flex flex-col items-center">
      <h2 class="font-bold dark:text-light-2 mb-2">{{ userInfo.nickname }}</h2>
      <p class="text-xs text-light-5 dark:text-light-2">
        {{ userInfo.introduction }}
      </p>
    </div>
    <div>
      <ul class="flex items-center border-t px-5">
        <li
          v-for="(item, index) in 3"
          class="h-12 flex items-center mx-5 cursor-pointer"
          :class="{
            'border-b-2': index === 0,
            'border-blue-500': index === 0,
          }"
        >
          Tab{{ index }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { useGetUserInfo } from "~/api/server";
const route = useRoute();
let userInfo = ref({});
async function init() {
  try {
    const { data } = await useGetUserInfo(route.params.id);
    userInfo.value = data;
    console.log(data);
  } catch (error) {
    console.log(error);
  }
}
init();
</script>
