<template>
  <div class="">
    <div class="grid grid-cols-9">
      <div class="h-full m-1 col-span-9 sm:col-span-7">
        <Card>
          <Icon name="uil:github" color="black" />
          <div class="text-xl font-bold">发帖子</div>
          <div class="border-t mt-5 py-5">
            <ul class="mb-5">
              <li v-for="item in getGroup">
                <KButton class="mr-5">{{ item.name }}</KButton>
              </li>
            </ul>
            <div class="flex items-center mb-5">
              <Input
                type="text"
                class="w-full rounded-lg p-2 mr-5"
                placeholder="请输入标题"
                v-model="formData.title"
              />
              <KButton class="flex-shrink-0" @click="uploadCover = !uploadCover">
                选择封面
              </KButton>
            </div>
            <!-- 默认编辑器 -->
            <Editor ref="content" v-model="formData.defaultPost.content" class="mb-5 mix-h-96" />
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
            <Editor class="mt-5" v-if="hiddenContentStatus" v-model="formData.defaultPost.hiddenContent" />
          </div>
          <!-- 附件 -->
          <div>
            <div class="flex mb-5">
              <Captcha class="w-32 shadow-center" />
              <Input
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
  <Dialog v-model="showUpload"></Dialog>

  <!-- 选择封面 -->
  <Dialog v-model="uploadCover">
    {{ activeList }}
    <div
      class="w-[800px] h-96 grid grid-cols-12 gap-2 px-5 m-5 overflow-y-auto"
    >
      <img
        @click="actived(item)"
        class="w-32 h-32 object-cover rounded-xl shadow-center col-span-2 hover:border"
        :class="{
          'shadow-active': activeList.includes(item),
        }"
        :src="item"
        v-for="item in [...$refs.content.getAllImages(), ...uploadList]"
      />
      <div
        class="w-32 h-32 col-span-2 flex items-center justify-center rounded-xl shadow-center active:shadow-active"
      >
        <Upload
          class="flex items-center justify-center"
          @uploadSuccess="(val) => uploadList.push(val)"
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
import { createPost } from "~/api";
let store = useGroupStore();
let userStore = useUserStore();
let userInfo = computed(() => userStore.getUserInfo);
let getGroup = computed(() => store.getGroup);
let showUpload = ref(false);
let uploadCover = ref(false);
let uploadList = ref([]);
let activeList = ref([]);
let hiddenContentStatus = ref(false);
let formData = ref({
  userId: userInfo.value.userId,
  title: "",
  tags: [],
  type: 1,
  cover: [],
  defaultPost: {
    content: "",
    hidden: hiddenContentStatus.value ? 1 : 0,
    hiddenContent: "",
  },
});
const { formatNumber } = useFormatNumber();
const { isMobile } = useMobileDetect();

definePageMeta({
  layout: "user",
});
onMounted(async () => {
  await store.fetchGroup();
});
const actived = (url) => {
  if (activeList.value.includes(url)) {
    activeList.value = activeList.value.filter((item) => item !== url);
  } else {
    activeList.value.push(url);
  }
  formData.value.cover = activeList.value;
};

const submit = async () => {
  console.log(formData.value);
  return
  try {
    let res = await createPost({
      title: "标题",
      content: $refs.content.getContent(),
      group_id: 1,
      cover: activeList.value,
    });
  } catch (error) {}
  console.log("submit");
};
</script>
