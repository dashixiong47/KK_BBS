<template>
  <div class="w-full sticky top-0 z-50">
    <ul class="w-10 absolute top-5 right-0 sm:-left-16">
      <li
        v-for="item in buttonList"
        @click="handleClick(item)"
        class="rounded-full cursor-pointer glass w-10 h-10 mb-5 flex items-center justify-center"
      >
        <Icon
          :name="item.icon"
          class="text-[--secondary-text] dark:text-[--dark-secondary-text]"
          :class="{
            '!text-[--color-primary] dark:!text-[--dark-color-primary]':
              getActive(item.type),
          }"
        />
        <div
          v-if="item.type"
          class="min-w-[16px] h-4 px-1 absolute -top-2 left-8 rounded-full bg-slate-400 text-white text-xs flex items-center justify-center"
        >
          {{ getData(item.type) }}
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { useUserStore, useLoginStore } from "~/stores/main";
import { topicLike, topicCollect } from "~/api";
const { formatNumber } = useFormatNumber();
let userInfo = useUserStore();
let loginStore = useLoginStore();
let isLogin = computed(() => userInfo.isLogin);
let props = defineProps({
  handleClick: {
    type: Function,
    default: () => {},
  },
  detail: {
    type: Object,
    default: () => {},
  },
});
const buttonList = ref([
  {
    name: "浏览",
    icon: "tabler:eye",
    type: "view",
  },
  {
    name: "点赞",
    icon: "icon-park-solid:thumbs-up",
    type: "like",
  },
  {
    name: "评论",
    icon: "mdi:comment-processing-outline",
    type: "comment",
  },
  {
    name: "收藏",
    icon: "tabler:bookmark",
    type: "collect",
  },
  {
    name: "分享",
    icon: "tabler:share",
  },
  // {
  //   name: "举报",
  //   icon: "tabler:report",
  // },
  // {
  //   name: "打赏",
  //   icon: "tabler:bookmark",
  // },
]);
function getData(type) {
  switch (type) {
    case "view":
      return formatNumber(props.detail.view);
    case "like":
      return formatNumber(props.detail.like);
    case "comment":
      return formatNumber(props.detail.comment);
    case "collect":
      return formatNumber(props.detail.collect);
    default:
      break;
  }
}
function getActive(type) {
  switch (type) {
    case "view":
      return true;
    case "like":
      return props.detail.likeState;
    case "collect":
      return props.detail.collectState;
    case "comment":
      return props.detail.commentState;
    default:
      return false;
  }
}
function handleClick(item) {
  switch (item.type) {
    case "comment":
      // 获取元素
      var element = document.getElementById("comment");
      // 滚动到元素
      element.scrollIntoView({ behavior: "smooth" });
      break;
    case "like":
      handleLike();
      break;
    case "collect":
      handleCollect();
      break;
    default:
      break;
  }
}
/**
 * @description 点赞
 */
const handleLike = async () => {
  if (!isLogin.value) {
    loginStore.setLoginStatus();
    return;
  }
  try {
    await topicLike(props.detail.id);

    props.detail.likeState = !props.detail.likeState;
    props.detail.likeState ? props.detail.like++ : props.detail.like--;
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
    await topicCollect(props.detail.id);
    props.detail.collectState = !props.detail.collectState;
    props.detail.collectState ? props.detail.collect++ : props.detail.collect--;
  } catch (error) {
    console.log(error);
  }
};
</script>
