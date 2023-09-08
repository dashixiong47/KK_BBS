<template>
  <div class="">
    <div class="grid grid-cols-9">
      <div class="h-full m-1 col-span-9 sm:col-span-7">
        <Card>
          <div class="text-xl font-bold">发帖子</div>
          <div class="border-t mt-5 py-5">
            <ul class="mb-5">
              <li v-for="item in getGroup">
                <KButton @click="formData.groupId=item.id"  class="mr-5">{{ item.name }}</KButton>
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
              >
                选择封面
              </KButton>
            </div>
            <!-- 默认编辑器 -->
            <Editor
              ref="content"
              v-model:moduleValue="formData.defaultPost.content"
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
              <KButton @click="hiddenContentStatus = !hiddenContentStatus">
                评论可见
              </KButton>
            </div>
            <!-- 隐藏内容 -->
            <Editor
              class="mt-5"
              v-if="hiddenContentStatus"
              v-model:moduleValue="formData.defaultPost.hiddenContent"
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
      <img
        @click="actived(item)"
        class="w-32 h-32 object-cover rounded-xl shadow-center col-span-2 hover:border"
        :class="{
          'shadow-active': activeCoverList.includes(item),
          border: activeCoverList.includes(item),
        }"
        :src="item"
        v-for="item in [...$refs.content.getAllImages(), ...uploadImgList]"
      />
      <div
        class="w-32 h-32 col-span-2 flex items-center justify-center rounded-xl shadow-center active:shadow-active"
      >
        <Upload
          class="flex items-center justify-center"
          @uploadSuccess="({url}) => uploadImgList.push(url)"
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
import Sortable from 'sortablejs';
import { useUserStore } from "~/stores/main.js";
import { useGroupStore } from "~/stores/init.js";
import { createPost } from "~/api";
let { addMessage } = useMessage();
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
  cover: [],
  defaultPost: {
    content: "",
    hidden: hiddenContentStatus.value ? 1 : 0,
    hiddenContent: "",
  },
  code: "",
  groupId: "",
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
  formData.value.cover = activeCoverList.value;
};

const submit = async () => {
  console.log(formData.value);
  let message = checkParams();
  if (message) {
    addMessage(message, "warning");
    return;
  }
  try {
    let res = await createPost(formData.value);
  } catch (error) {
    captchaRef.value.init();
    addMessage(error, "error");
  }
  console.log("submit");
};
// 校验参数
const checkParams = () => {
  if (!formData.value.title) {
    return "请输入标题";
  }
  if (!formData.value.defaultPost.content) {
    return "请输入内容";
  }
  if (!formData.value.groupId) {
    return "请选择分组";
  }
  if (!formData.value.code) {
    return "请先完成验证";
  }
  return false;
};
</script>
