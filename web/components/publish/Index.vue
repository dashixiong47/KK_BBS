<template>
  <div>
    
    <!-- 默认编辑器 -->
    <Editor
      ref="content"
      v-model:moduleValue="formData.topicBasic.content"
      placeholder="请输入"
      class="mb-5 mix-h-96"
    />

    <!-- 是否回复可见 -->
    <p class="flex">评论可见:<KSwitch class="ml-2" v-model="hiddenContentStatus"></KSwitch></p>

    <!-- 隐藏内容 -->
    <Editor
      class="mt-5"
      v-if="hiddenContentStatus"
      v-model:moduleValue="formData.topicBasic.hiddenContent"
    />
  </div>
</template>

<script setup>
// 隐藏内容状态
let hiddenContentStatus = ref(false);

let formData = ref({
  topicBasic: {
    content: "",
    hidden: hiddenContentStatus.value ? 1 : 0,
    hiddenContent: "",
  },
});
// 校验参数
const checkParams = () => {
  if (!formData.value.topicBasic.content) {
    return t("topic-create-content");
  }
  return false;
};
// 提交
const getFormData = () => {
  return formData.value;
};
defineExpose({
  getFormData,
  checkParams,
});
</script>
