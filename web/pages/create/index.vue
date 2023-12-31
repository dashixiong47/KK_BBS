<template>
  <div class="">
    <div class="grid grid-cols-9">
      <div class="h-full m-1 col-span-9">
        <Card>
          <!-- 帖子类型 -->
          <TopicType @selectd="selectdTopicType" />
          <!-- 通用 标题封面 -->
          <div class="border-t mt-5 py-5">
            <Skeleton :lines="3" v-if="!getGroup.length" />
            <ul class="mb-5 flex">
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
              <!-- 选择封面 -->
              <UploadCover @selectd="(val) => (formData.covers = val)" />
            </div>
            <component ref="publish" :is="selectdTopic"></component>
          </div>
          <!-- 是否回复可见 -->
          <p class="flex mb-5">
            上传附件:<KSwitch class="ml-2" v-model="uploadShow"></KSwitch>
          </p>

          <!-- 上传附件 -->
          <Attachment v-if="uploadShow" v-model="fileList"></Attachment>

          <div class="mt-5">
            <div class="flex mb-5">
              <Captcha
                ref="captchaRef"
                class="w-32 shadow-center"
                @captchaId="(id) => (formData.captchaId = id)"
              />
              <KInput
                v-model="formData.code"
                type="text"
                class="w-36 rounded-lg p-2 ml-5"
                placeholder="验证码"
              />
            </div>
            <KButton
              class="glass px-5 py-2 w-full font-main-color rounded-lg overflow-hidden"
              @click="submit"
            >
              发布
            </KButton>
          </div>
        </Card>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useGroupStore, useAppConfigStore } from "~/stores/init.js";
import { createTopic } from "~/api";
import { useUserStore } from "~/stores/main.js";
import Topic from "~/components/publish/Index.vue";
import Video from "~/components/publish/Video.vue";
import Text from "~/components/publish/Text.vue";
import Image from "~/components/publish/Image.vue";
useHead({
  title: "发布",
});
const { t } = useI18n();
const appConfigStore = useAppConfigStore();
let { addMessage } = useMessage();
let { to } = useToRoute();
let store = useGroupStore();
let getGroup = computed(() => store.getGroup.filter((item) => item.id > 0));
let captchaRef = ref(null);
let TopicTypeOptions = ref({
  topic: Topic,
  video: Video,
  text: Text,
  image: Image,
});
// 上传文件显示of隐藏
let uploadShow = ref(false);
// 发布组件
let publish = ref(null);

// 附件列表
let fileList = ref([]);
// 表单数据
let formData = ref({
  code: "",
  title: "",
  tags: [],
  covers: [],
  type: 0,
  groupId: getGroup.value[0]?.id,
  captchaId: "",
  attachment: fileList.value,
});
let type = null;
let selectdTopic = ref(TopicTypeOptions.value.topic);
const selectdTopicType = (item) => {
  formData.value.type = item.index;
  selectdTopic.value = TopicTypeOptions.value[item.type];
  type = item.type;
};
const submit = async () => {
  let message = checkParams();
  if (message) {
    addMessage(message, "warning");
    return;
  }
  let form = {
    ...formData.value,
    ...publish.value.getFormData(),
  };
  let regex = new RegExp(`${appConfigStore.host}`, "g");

  try {
    let { data } = await createTopic(
      JSON.parse(JSON.stringify(form).replace(regex, ""))
    );
    to(`/${type}/${data}`);
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
  if (!formData.value.groupId) {
    return t("topic-create-groupId");
  }
  if (!formData.value.code) {
    return t("topic-create-code");
  }
  let state = publish.value.checkParams();
  if (state) {
    return state;
  }
  return false;
};
// 获取小组列表
store.fetchGroup();
</script>
