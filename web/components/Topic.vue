<template>
  <div class="flex mb-5">
    <Avatar :url="detail.user?.avatar" />
    <div class="w-full ml-2 flex items-center justify-between">
      <div class="flex flex-col items-start">
        <KLink to="#" class="text-md font-bold primary-text">
          {{ detail.user?.nickname }}
        </KLink>
        <span class="text-xs secondary-text">
          发布时间:{{ getRelativeTime(detail.createdAt) }}
        </span>
      </div>
      <div>
        <span class="mr-5 regular-text text-xs">
          <Icon name="tabler:eye" class="mr-1" />{{ formatNumber(detail.view) }}
        </span>
        <span class="regular-text text-xs">
          <Icon name="ic:round-location-on" class="mr-1" />上海
        </span>
      </div>
    </div>
  </div>
  <KLink :to="`/topic/${detail.id}`" class="">
    <p class="text-md mb-3 primary-text">{{ detail.title }}</p>
    <div class="regular-text text-sm mb-5">
      {{ detail.summary }}
    </div>
    <!-- <ul class="flex flex-row flex-wrap "> -->
    <ul class="grid grid-cols-6">
      <li
        v-for="(item, index) in images"
        class="shadow-center rounded-2xl pt-[66.67%] relative flex items-center justify-center m-2 overflow-hidden"
        :class="getClassName(images.length, index)"
      >
        <KImage
          class="object-cover absolute inset-0"
          :source="getPath(item)"
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
        class="mr-5 flex-shrink-0 rounded-2xl border bg-light-2 px-2 text-sm flex items-center secondary-text"
      >
        <KLink to="#"> #分类{{ index }} </KLink>
      </li>
    </template>
  </ul>
  <div class="my-2 grid grid-cols-12">
    <button
      v-for="btn in btns"
      class="regular-text col-span-3 py-2 flex items-center justify-center rounded-lg font-medium text-sm hover:shadow-default"
      :class="[
        btn.active
          ? '!text-[--color-primary] dark:!text-[--dark-color-primary]'
          : '',
      ]"
      @click="btn.func()"
    >
      <Icon :name="btn.icon" size="1rem" />
      <span class="mx-1">{{ btn.count }}</span>
      <span v-if="!isMobile"> {{ btn.text }}</span>
    </button>
  </div>
</template>
<script setup>
import { ref, onMounted } from "vue";
import useMobileDetect from "~/composables/useMobileDetect";
import useFormatNumber from "~/composables/useFormatNumber";
import { useUserStore, useLoginStore } from "~/stores/main";

import { topicLike, topicCollect } from "~/api";
const { to } = useToRoute();
let userInfo = useUserStore();
let loginStore = useLoginStore();
let isLogin = computed(() => userInfo.isLogin);
const { getRelativeTime } = useTime();
const { getPath } = usePath();
let { detail } = defineProps({
  // 详情
  detail: {
    type: Object,
    default: () => ({}),
  },
});
const btns = reactive([
  {
    icon: "icon-park-solid:thumbs-up",
    text: "点赞",
    count: detail.like,
    active: detail.likeState,
    func: () => handleLike(),
  },
  {
    icon: "icon-park-solid:message-one",
    text: "评论",
    count: detail.comment,
    active: detail.commentState,
    func: () => {
      to(`/topic/${detail.id}`);
    },
  },
  {
    icon: "icon-park-solid:collection-files",
    text: "收藏",
    count: detail.collect,
    active: detail.collectState,
    func: () => handleCollect(),
  },
  {
    icon: "tabler:share",
    text: "分享",
    count: detail.like,
    func: () => {},
  },
]);
const listContainer = ref(null);
const items = ref(Array(3).fill(null));
const itemsToShow = ref([...items.value]);
let images = computed(() => {
  return detail.covers;
});
const { formatNumber } = useFormatNumber();
const { isMobile } = useMobileDetect();

/**
 * @argument {number} length
 * @argument {number} index
 * @returns {string}
 * @description 根据图片数量和索引返回对应的样式
 */
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
/**
 * @description 点赞
 */
const handleLike = async () => {
  if (!isLogin.value) {
    loginStore.setLoginStatus();
    return;
  }
  try {
    await topicLike(detail.id);
    btns[0].active = !btns[0].active;
    btns[0].active ? btns[0].count++ : btns[0].count--;
  } catch (error) {
    console.log(error);
  }
};
// 收藏
const handleCollect = async () => {
  if (!isLogin.value) {
    loginStore.setLoginStatus();
    return;
  }
  try {
    await topicCollect(detail.id);
    btns[2].active = !btns[2].active;
    btns[2].active ? btns[2].count++ : btns[2].count--;
  } catch (error) {
    console.log(error);
  }
};
</script>
