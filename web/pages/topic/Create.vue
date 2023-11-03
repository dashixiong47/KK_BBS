<template>
  <div class="">
    <div class="grid grid-cols-9">
      <div class="h-full m-1 col-span-9">
        <Card>
          <!-- 帖子类型 -->
          <TopicType />
          <!-- 通用 标题封面 -->
          <div class="border-t mt-5 py-5">
            <Skeleton :lines="3" v-if="!getGroup.length" />
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
            <component ref="publish" :is="Topic"></component>
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
              class="glass px-5 py-2 w-full primary-text rounded-lg overflow-hidden"
              @click="submit"
            >
              发布
            </KButton>
          </div>
        </Card>
      </div>
    </div>
  </div>

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
import { useGroupStore } from "~/stores/init.js";
import { createTopic } from "~/api";
import { useUserStore } from "~/stores/main.js";
import Topic from "~/components/publish/Index.vue";
let userStore = useUserStore();
const { t } = useI18n();

let { addMessage } = useMessage();
let { to } = useToRoute();
let store = useGroupStore();
let getGroup = computed(() => store.getGroup);
let captchaRef = ref(null);
// 上传文件显示of隐藏
let uploadShow = ref(false);
// 发布组件
let publish = ref(null);
// 用户信息
let userInfo = computed(() => userStore.getUserInfo);
let uploadCover = ref(false);
let uploadImgList = ref([]);
// 附件列表
let fileList = ref([]);
let activeCoverList = ref([]);
// 表单数据
let formData = ref({
  code: "",
  title: "",
  tags: [],
  covers: [],
  type: 1,
  userId: userInfo.value.id,
  groupId: getGroup.value[0]?.id,
  captchaId: "",
  attachment: fileList.value,
});
// 使用layout
definePageMeta({
  layout: "user",
});
// 获取小组列表
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
  let message = checkParams();
  if (message) {
    addMessage(message, "warning");
    return;
  }
  try {
    let data = await createTopic({
      ...formData.value,
      ...publish.value.getFormData(),
    });
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
  if (!formData.value.groupId) {
    return t("topic-create-groupId");
  }
  if (!formData.value.code) {
    return t("topic-create-code");
  }
  return false;
};
</script>
