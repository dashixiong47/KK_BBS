<template>
  <p class="mb-2 w-full">评论</p>
  <ReplyInput v-model="replyData" @success="replySuccess" />
  <ReplyInput
    class="mt-2"
    v-show="showSubReply"
    ref="subReplyIpt"
    v-model="subReplyData"
    @success="replySuccess"
  />
  <div class="flex items-center mb-5">
    <div class="flex items-center">
      <span class="text-base font-extrabold">{{ $t("comments-title") }}</span>
      <span class="ml-2 text-sm font-normal text-light-5 font-regular-color">{{
        total
      }}</span>
    </div>
    <div class="ml-5 text-sm text-light-5">
      <button
        class="border px-2 rounded-l-md"
        :class="[
          !selected
            ? 'text-[--illuminate-color] dark:text-[--dark-illuminate-color]'
            : '',
        ]"
        @click="change(0)"
      >
        {{ $t("comments-hot") }}
      </button>
      <button
        class="border px-2 rounded-r-md"
        :class="{
          'bg-light-3 dark:bg-light-4': selected,
          'text-blue-400': selected,
        }"
        @click="change(1)"
      >
        {{ $t("comments-new") }}
      </button>
    </div>
  </div>
  <ul>
    <li
      v-for="(item, index) in comments"
      class="flex mb-5 pb-5"
      :class="{
        'border-b': index !== comments.length - 1,
      }"
    >
      <Avatar :url="item.user?.avatar" />
      <div class="ml-5 w-full">
        <p class="mb-2"><KName :name="item.user?.nickname" /></p>
        <div class="text-sm">
          {{ item.content }}
        </div>
        <div
          class="text-xs text-light-5 flex items-center text-[--font-secondary-color] dark:text-[--dark-font-secondary-color]"
        >
          <span class="mr-2 cursor-pointer">
            {{ getRelativeTime(item.createdAt) }}
          </span>
          <span
            class="mr-2 cursor-pointer flex items-center"
            :class="[
              item.likeState
                ? 'text-[--illuminate-color] dark:text-[--dark-illuminate-color]'
                : '',
            ]"
            @click="commentLikeChange(item, item.id, 0)"
          >
            <Icon name="icon-park-solid:thumbs-up" class="mr-1" />
            {{ item.like }}
          </span>
          <span
            class="cursor-pointer text-[--illuminate-color] dark:text-[--dark-illuminate-color]"
            @click="(e) => subReply(e, item, item.id)"
            >{{ $t("comments-reply") }}
          </span>
        </div>
        <div class="reply"></div>
        <div>
          <SubComments
            :total="item.total"
            :parentId="item.id"
            :topicId="topicId"
            :reply="item.reply"
            :subReply="subReply"
            :commentLikeChange="commentLikeChange"
          />
        </div>
      </div>
    </li>
  </ul>
  <div>
    <KButton
      class="w-full flex justify-center"
      :disabled="getDisabled"
      :throttle="loadMore"
      :class="{ 'cursor-not-allowed': getDisabled }"
      >{{ getDisabled ? "没有更多了" : "加载更多" }}</KButton
    >
  </div>
</template>

<script setup>
import { useGetComments } from "~/api/server";
import { commentLike } from "~/api";
import { useUserStore, useLoginStore } from "~/stores/main";
let { topicId } = defineProps({
  topicId: {
    type: String,
    required: true,
  },
});
const { t } = useI18n();
const { notice } = useNotice();
const { getRelativeTime } = useTime();
let userInfo = useUserStore();
let loginStore = useLoginStore();
let isLogin = computed(() => userInfo.isLogin);

let lastReplyNode = null;
let showSubReply = ref(false);
let comments = ref([]);
let total = ref(0);
let replyData = ref({
  topicId: topicId,
  content: null,
});
let subReplyData = ref({
  topicId: topicId,
  parentId: null,
  replyToUserId: null,
  content: null,
});
let type = ref("all");
let page = {
  page: 1,
  pageSize: 10,
};
let emit = defineEmits(["change"]);
let selected = ref(0);
let subReplyIpt = ref(null);
const getDisabled = computed(() => {
  return comments.value.length === total.value;
});
const change = (num) => {
  if (selected.value === num) return;
  selected.value = num;
  switch (num) {
    case 0:
      type.value = "all";
      break;
    case 1:
      type.value = "hot";
      break;
    default:
      type.value = "all";
      break;
  }
  page = {
    page: 1,
    pageSize: 10,
  };
  comments.value = [];
  getComments();
};
const replySuccess = () => {
  showSubReply.value = false;
  resetLastReplyNode();
  getComments();
  emit("change");
};

// ----------------------------------------------------------------
// 辅助函数：切换按钮文本
const toggleButtonText = (target, value) => {
  target.innerText = value ? t("cancel-comments-reply") : t("comments-reply");
};

// 辅助函数：重置 lastReplyNode 变量
const resetLastReplyNode = () => {
  if (lastReplyNode) {
    toggleButtonText(lastReplyNode.target, false); // 重置按钮文本
    lastReplyNode = null; // 清空 lastReplyNode
  }
};

// 主函数：处理回复按钮的点击事件
const subReply = (e, item, parentId) => {
  const { target } = e; // 获取事件目标
  const replyNode = target.parentNode.parentNode.querySelector(".reply"); // 获取回复节点
  const subReplyNode = subReplyIpt.value.$el; // 获取子回复节点

  if (lastReplyNode) {
    // 如果存在上一个回复节点
    const isSameTarget = lastReplyNode.target === target; // 检查是否是同一个目标
    showSubReply.value = isSameTarget ? !showSubReply.value : true; // 根据条件切换 showSubReply 的值
    toggleButtonText(target, showSubReply.value); // 更新按钮文本
    if (!isSameTarget) {
      // 如果不是同一个目标
      resetLastReplyNode(); // 重置上一个回复节点
    }
  } else {
    // 如果不存在上一个回复节点
    showSubReply.value = true; // 设置 showSubReply 为 true
    toggleButtonText(target, true); // 更新按钮文本
  }

  subReplyData.value.parentId = parentId; // 设置 parentId
  subReplyData.value.replyToUserId = item.user.id; // 设置 replyToUserId
  replyNode.appendChild(subReplyNode); // 将子回复节点添加到回复节点
  lastReplyNode = e; // 更新 lastReplyNode 为当前事件
};
// ----------------------------------------------------------------

const loadMore = () => {
  console.log("----");
  if (page.page * page.pageSize >= total.value) {
    notice({
      title: "提示",
      content: "没有更多了",
    });
    return;
  }
  page.page++;
  getComments();
};
// 获取评论
async function getComments() {
  try {
    let { data } = await useGetComments(topicId, { ...page, type: type.value });
    // 计算当前页面在 comments.value 数组中的起始和结束索引
    const startIndex = (page.page - 1) * page.pageSize;
    const endIndex = startIndex + page.pageSize;
    // 如果 comments.value 有数据，则替换；否则，连接。
    if (comments.value.length > 0) {
      // 删除旧的数据范围，并插入新的数据
      comments.value.splice(
        startIndex,
        endIndex - startIndex,
        ...data.comments
      );
    } else {
      // 如果 comments.value 是空的，直接连接新数据
      comments.value = comments.value.concat(data.comments);
    }
    total.value = data.total;
  } catch (error) {
    console.log(error);
  }
}
// 评论点赞
async function commentLikeChange(item, commentId, subCommentId) {
  if (!isLogin.value) {
    loginStore.setLoginStatus();
    return;
  }
  let options = {
    commentId,
  };
  if (subCommentId) {
    options["subCommentId"] = subCommentId;
  }
  try {
    let data = await commentLike(topicId, options);

    item.likeState = !item.likeState;

    if (item.likeState) {
      item.like += 1;
    } else {
      item.like -= 1;
    }
  } catch (error) {}
}

function init() {
  getComments();
}
init();
</script>
