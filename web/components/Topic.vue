<template>
  <div class="flex mb-5">
    <Avatar />
    <div class="w-full ml-2 flex items-center justify-between">
      <div class="flex flex-col items-start">
        <KLink to="#" class="text-md font-bold"> 名称 </KLink>
        <span class="text-xs text-gray-400">发布时间</span>
      </div>
      <div>
        <span class="mr-5 text-light-5 text-xs">
          <Icon name="icon-kejian" class="mr-1" />{{ formatNumber(11111) }}
        </span>
        <span class="text-light-5 text-xs">
          <Icon name="icon-shouye" class="mr-1" />上海
        </span>
      </div>
    </div>
  </div>
  <KLink to="/topic/1" class="">
    <p class="text-md mb-3">这里是标题</p>
    <div class="text-gray-400 text-sm dark:text-dark-6">
      Lorem ipsum dolor sit amet,consectetur adipiscing elit. Morbi nulla
      dolor,ornare at commodo non,feugiat non nisi. Phasellus faucibus mollis
      pharetra. Proin blandit ac massa sed rhoncus See more
    </div>
    <!-- <ul class="flex flex-row flex-wrap "> -->
    <ul class="grid grid-cols-6">
      <li
        v-for="(item, index) in images.slice(0, 9)"
        class="pt-[66.67%] relative flex items-center justify-center m-2 overflow-hidden"
        :class="getClassName(images.length, index)"
      >
        <Image
          class="rounded-2xl object-cover absolute inset-0"
          :source="item.url"
          alt=""
          srcset=""
        />
      </li>
    </ul>
  </KLink>
  <ul
    class="w-full h-6 flex overflow-x-auto"
    :class="{
      'hidden-scrollbar': isMobile,
    }"
    ref="listContainer"
  >
    <template v-for="(item, index) in itemsToShow">
      <li
        class="mr-5 flex-shrink-0 rounded-2xl border bg-light-2 px-2 text-sm flex items-center text-light-5 dark:text-dark-2"
      >
        <KLink to="#"> #分类{{ index }} </KLink>
      </li>
    </template>
  </ul>
  <div class="my-2 grid grid-cols-12">
    <button
      class="col-span-3 py-2 flex items-center justify-center rounded-lg font-medium text-sm text-dark-1 dark:text-light-2 hover:bg-light-3 dark:hover:bg-light-4 dark:hover:text-dark-2"
      :class="{
        '!text-blue-400': true,
      }"
    >
      <Icon name="icon-dianzan"> </Icon>
      <span class="mx-1">{{ formatNumber(1000) }}</span>
      <span v-if="!isMobile"> 点赞</span>
    </button>
    <button
      class="col-span-3 py-2 flex items-center justify-center rounded-lg font-medium text-sm text-dark-1 dark:text-light-2 hover:bg-light-3 dark:hover:bg-light-4 dark:hover:text-dark-2"
    >
      <Icon name="icon-pinglun"></Icon> <span class="mx-1">123</span>
      <span v-if="!isMobile">评论</span>
    </button>
    <button
      class="col-span-3 py-2 flex items-center justify-center rounded-lg font-medium text-sm text-dark-1 dark:text-light-2 hover:bg-light-3 dark:hover:bg-light-4 dark:hover:text-dark-2"
    >
      <Icon name="icon-dianzan"></Icon> <span class="mx-1">123</span>
      <span v-if="!isMobile"> 点赞</span>
    </button>
    <button
      class="col-span-3 py-2 flex items-center justify-center rounded-lg font-medium text-sm text-dark-1 dark:text-light-2 hover:bg-light-3 dark:hover:bg-light-4 dark:hover:text-dark-2"
    >
      <Icon name="icon-shoucang"></Icon> <span class="mx-1">123</span>
      <span v-if="!isMobile"> 收藏</span>
    </button>
  </div>
</template>
<script setup>
import { ref, onMounted } from "vue";
const images = [
  {
    url: "https://images.hxsj.in/test/1.png",
  },
  {
    url: "https://images.hxsj.in/test/2.png",
  },
  {
    url: "https://images.hxsj.in/test/1.png",
  },
  {
    url: "https://images.hxsj.in/test/2.png",
  },
  {
    url: "https://images.hxsj.in/test/1.png",
  },
  {
    url: "https://images.hxsj.in/test/2.png",
  },
  {
    url: "https://images.hxsj.in/test/2.png",
  },
  {
    url: "https://images.hxsj.in/test/1.png",
  },
  {
    url: "https://images.hxsj.in/test/2.png",
  },
];
const listContainer = ref(null);
const items = ref(Array(3).fill(null));
const itemsToShow = ref([...items.value]);
import useMobileDetect from "~/composables/useMobileDetect";
import useFormatNumber from "~/composables/useFormatNumber";
const { formatNumber } = useFormatNumber();
const { isMobile } = useMobileDetect();
const getClassName = (length, index) => {
  switch (length) {
    case 1:
    case 2:
    case 4:
      return "col-span-3";
    case 5:
    case 8:
      if (index < 2) {
        return "col-span-3";
      }
      return "col-span-2";
    case 7:
      if (index < 4) {
        return "col-span-3";
      }
      return "col-span-2";
    default:
      return "col-span-2";
  }
};
</script>
