<template>
  <div class="flex mb-5">
    <Avatar />
    <div class="ml-2">
      <div class="flex flex-col items-start">
        <NuxtLink to="#" class="text-md font-bold"> 名称 </NuxtLink>
        <span class="text-xs text-gray-400">发布时间</span>
      </div>
    </div>
  </div>
  <NuxtLink to="#" class="">
    <p class="text-md mb-3">这里是标题</p>
    <div class="text-gray-600 text-sm">测试 测试 测试</div>
    <ul>
      <li v-for="(item, index) in images">
        <Image
          :class="getClassName(images.length, index)"
          :source="item.url"
          alt=""
          srcset=""
        />
      </li>
    </ul>
  </NuxtLink>
</template>
<script setup>
const images = [
  // {
  //   url: "https://file.mlog.club/mlog.club/logo_364x200.png",
  // },
  // {
  //   url: "https://file.mlog.club/mlog.club/logo_364x200.png",
  // },
  // {
  //   url: "https://file.mlog.club/mlog.club/logo_364x200.png",
  // },
  // {
  //   url: "https://file.mlog.club/mlog.club/logo_364x200.png",
  // },
  // {
  //   url: "https://file.mlog.club/mlog.club/logo_364x200.png",
  // },
  // {
  //   url: "https://file.mlog.club/mlog.club/logo_364x200.png",
  // },
  // {
  //   url: "https://file.mlog.club/mlog.club/logo_364x200.png",
  // },
  // {
  //   url: "https://file.mlog.club/mlog.club/logo_364x200.png",
  // },
  // {
  //   url: "https://file.mlog.club/images/2023/07/18/06b529f47842fd717f74f310916b36fb.png!preview",
  // },
  {
    url: "https://file.mlog.club/images/2023/08/04/a7d614e2a237c71a4863be8cdd033f74.jfif!avatar",
  },
];
const getClassName = (length, index) => {
  switch (length) {
    case 1:
      return "w-full h-16";

    default:
      break;
  }
};
</script>
