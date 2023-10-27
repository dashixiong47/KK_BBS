<template>
  <div>
    <!-- 默认编辑器 -->
    <Editor
      ref="content"
      v-model:moduleValue="formData.topicBasic.content"
      placeholder="请输入"
      class="mb-5 mix-h-96"
    />
    <!-- 上传附件按钮 -->
    <div class="mb-5">
      <KButton class="mr-5" @click="showUpload = !showUpload">
        上传附件
      </KButton>
    </div>
    <!-- 是否回复可见 -->
    <div class="flex">
      <KButton
        @click="hiddenContentStatus = !hiddenContentStatus"
        :actived="hiddenContentStatus"
      >
        评论可见
      </KButton>
    </div>
    <!-- 隐藏内容 -->
    <Editor
      class="mt-5"
      v-if="hiddenContentStatus"
      v-model:moduleValue="formData.topicBasic.hiddenContent"
    />
  </div>
</template>

<script setup>

import { useUserStore } from "~/stores/main.js";
let userStore = useUserStore();
// 用户信息
let userInfo = computed(() => userStore.getUserInfo);
// 隐藏内容状态
let hiddenContentStatus = ref(false);
// 上传附件状态
let showUpload = ref(false);
let formData = ref({
  userId: userInfo.value.id,
  title: "",
  tags: [],
  type: 1,
  covers: [],
  topicBasic: {
    content: "",
    hidden: hiddenContentStatus.value ? 1 : 0,
    hiddenContent: "",
  },

});
</script>
