<template>
  <div class="flex flex-col items-end mb-5">
    <textarea
      :value="props.modelValue.content"
      @input="outInput"
      class="shadow-center bg-[rgba(0,0,0,0)] w-full h-16 text-sm border-[--border-light] dark:border-[--dark-border-light] mb-2 rounded-lg focus:outline-none p-2"
      :placeholder="props.placeholder"
    />
    <KButton class="w-20 h-10 text-sm" @click="postCommentCreate">发表</KButton>
  </div>
</template>

<script setup>
import { commentCreate } from "~/api";
import { useUserStore, useLoginStore } from "~/stores/main";
const { notice } = useNotice();
let userInfo = useUserStore();
let loginStore = useLoginStore();
let isLogin = computed(() => userInfo.isLogin);
let props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({}),
  },
  placeholder: {
    type: String,
    default: "请输入内容",
  },
});

let emit = defineEmits(["update:modelValue", "success"]);
const outInput = (e) => {
  emit("update:modelValue", {
    ...props.modelValue,
    content: e.target.value,
  });
};
const postCommentCreate = async () => {
  if (!isLogin.value) {
    loginStore.setLoginStatus();
    return;
  }
  if (!props.modelValue.content) {
    notice({
      title: "错误",
      content: "请输入内容",
    });
    return;
  }
  try {
    const res = await commentCreate(props.modelValue);
    notice({
      title: "评论成功",
      content: "评论成功",
    });
    emit("update:modelValue", {
      topicId: props.modelValue.topicId,
      content: null,
    });
    emit("success");
  } catch (error) {
    notice({
      title: "错误",
      content: error,
    });
  }
};
</script>
