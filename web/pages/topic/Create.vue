<template>
  <div class="">
    <div class="grid grid-cols-9">
      <div class="h-full m-1 col-span-9 sm:col-span-7">
        <Card>
          <div class="text-xl font-bold">发帖子</div>
          <TopicType />
          <div class="border-t mt-5 py-5">
            <ul class="mb-5">
              <li v-for="item in getGroup">
                <KButton
                  @click="formData.groupId = item.id"
                  class="mr-5 flex items-center"
                  :actived="formData.groupId === item.id"
                >
                  {{ item.name }}
                </KButton>
              </li>
            </ul>
            <div class="flex items-center mb-5">
              <KInput
                type="text"
                class="w-full rounded-lg p-2 mr-5"
                placeholder="请输入标题"
                v-model="formData.title"
              />
              <KButton
                class="flex-shrink-0"
                @click="uploadCover = !uploadCover"
                :actived="activeCoverList.length > 0"
              >
                选择封面
              </KButton>
            </div>
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
              <KButton @click="hiddenContentStatus = !hiddenContentStatus" :actived="hiddenContentStatus">
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
          <!-- 附件 -->
          <div>
            <div class="flex mb-5">
              <Captcha
                ref="captchaRef"
                class="w-32 shadow-center"
                @captchaId="(id) => (formData.captchaId = id)"
              />
              <KInput
                v-model="formData.code"
                type="text"
                class="border w-36 rounded-lg p-2 ml-5"
                placeholder="验证码"
              />
            </div>
            <KButton
              class="glass px-5 py-2 w-full primary-text rounded-lg overflow-hidden"
              @click="submit"
            >
              发布
            </KButton>
          </div>
        </Card>
      </div>
      <div class="hidden sm:col-span-2 sm:block m-1">
        <Card class="sticky top-0">
          <p class="border-b pb-2 mb-2 font-bold">附件</p>
          <ul>
            <li>121212</li>
          </ul>
        </Card>
      </div>
    </div>
  </div>
  <!-- 上传附件 -->
  <Dialog v-model="showUpload">
    <div class="w-[800px] h-96">
      {{ uploadGattachmentList }}
      <DragColumn :items="uploadGattachmentList"></DragColumn>
      <Upload
        accept="*"
        multiple
        @uploadSuccess="(val) => uploadGattachmentList.push(val)"
      >
        <KButton>上传</KButton>
      </Upload>
    </div>
  </Dialog>

  <!-- 选择封面 -->
  <Dialog v-model="uploadCover">
    <div
      class="w-[800px] h-96 grid grid-cols-12 gap-2 px-5 m-5 overflow-y-auto"
    >
      <div
        v-for="item in [...$refs.content.getAllImages(), ...uploadImgList]"
        class="col-span-2 w-32 h-32 relative cursor-pointer"
      >
        <img
          @click="actived(item)"
          class="w-full h-full object-cover rounded-xl shadow-center hover:border"
          :src="item"
        />
        <div
          v-if="activeCoverList.includes(item)"
          class="w-full h-full pointer-events-none bg-[--masking-color] dark:bg-[--dark-masking-color] absolute top-0 left-0"
        >
          <Icon
            name="ep:check"
            size="2rem"
            class="text-[--color-white] dark:text-[--dark-color-white] absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2"
          />
        </div>
      </div>
      <div
        class="w-32 h-32 col-span-2 flex items-center justify-center rounded-xl shadow-center active:shadow-active"
      >
        <Upload
          class="flex items-center justify-center"
          @uploadSuccess="({ url }) => uploadImgList.push(url)"
        >
          +
        </Upload>
      </div>
    </div>
    <div class="flex justify-end">
      <KButton @click="uploadCover = !uploadCover" class="primary-text">
        确定
      </KButton>
    </div>
  </Dialog>
</template>

<script setup>

import { useUserStore } from "~/stores/main.js";
import { useGroupStore } from "~/stores/init.js";
import { createTopic } from "~/api";
const { t } = useI18n();
let { addMessage } = useMessage();
let { to } = useToRoute();
let store = useGroupStore();
let userStore = useUserStore();
let userInfo = computed(() => userStore.getUserInfo);
let getGroup = computed(() => store.getGroup);
let captchaRef = ref(null);
let showUpload = ref(false);
let uploadCover = ref(false);
let uploadImgList = ref([]);
let uploadGattachmentList = ref([]);
let activeCoverList = ref([]);
let hiddenContentStatus = ref(false);
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
  code: "",
  groupId: getGroup.value[0]?.id,
  captchaId: "",
});
definePageMeta({
  layout: "user",
});
store.fetchGroup();
const actived = (url) => {
  if (activeCoverList.value.includes(url)) {
    activeCoverList.value = activeCoverList.value.filter(
      (item) => item !== url
    );
  } else {
    activeCoverList.value.push(url);
  }
  formData.value.covers = activeCoverList.value;
};

const submit = async () => {
  console.log(formData.value);
  let message = checkParams();
  if (message) {
    addMessage(message, "warning");
    return;
  }
  try {
    let data = await createTopic(formData.value);
    to("/topic/" + data);
    addMessage(t("submitted-success"));
  } catch (error) {
    captchaRef.value.init();
    addMessage(error, "error");
  }
};
// 校验参数
const checkParams = () => {
  if (!formData.value.title) {
    return t("topic-create-title");
  }
  if (!formData.value.topicBasic.content) {
    return t("topic-create-content");
  }
  if (!formData.value.groupId) {
    return t("topic-create-groupId");
  }
  if (!formData.value.code) {
    return t("topic-create-code");
  }
  return false;
};
</script>
