<template>
  <Card v-if="attachmentList.length">
    <p class="border-b pb-2 mb-2 font-bold">附件</p>
    <ul>
      <li
        class="p-2 mb-2 rounded-xl flex justify-between items-center"
        v-for="item in attachmentList"
      >
        <div class="flex w-3/4">
          <div class="mr-2 max-w-[60%] truncate">{{ item.name }}</div>
          <div><Icon name="ri:coins-fill"></Icon>{{ item.coins }}</div>
        </div>
        <KButton
          v-if="!item.status"
          class="shadow-default rounded-xl"
          @click="change(item)"
        >
          <Icon name="ic:baseline-add-card"></Icon>
        </KButton>
        <a v-else :href="item.fileUrl" :download="item.name">
          <KButton>
            <Icon name="material-symbols:download-rounded"></Icon>
          </KButton>
        </a>
      </li>
    </ul>
  </Card>
  <Dialog
    v-model="showTips"
    cancel-btn
    confirm-text="购买"
    cancel-text="放弃"
    @confirm="buy(attachment)"
  >
    <div class="w-96">
      <div>名称:"{{ attachment.name }}"</div>
      <div>类型:"{{ attachment.name }}"</div>
      <div>金币:"{{ attachment.coins }}"</div>
      <div>备注:"{{ attachment.remake }}"</div>
    </div>
  </Dialog>
</template>

<script setup>
import { buyAttachment } from "~/api";
import { useGetAttachment } from "~/api/server";
import { useLoginStore, useUserStore } from "~/stores/main.js";
const store = useLoginStore();
const userStore = useUserStore();
const userInfo = computed(() => userStore.getUserInfo);
const isLogin = computed(() => userStore.getIsLogin);
const { addMessage } = useMessage();
let props = defineProps({
  topicId: {
    type: String,
    default: "",
  },
});
let showTips = ref(false);
let attachmentList = ref([]);
let attachment = ref({});
async function init() {
  try {
    let { data } = await useGetAttachment(props.topicId);
    attachmentList.value = data;
  } catch (error) {}
}
function change(item) {
  if (!isLogin.value) {
    store.setLoginStatus();
    return;
  }
  attachment.value = item;
  showTips.value = true;
}
async function buy(item) {
  console.log(item);
  if (item.status) {
    addMessage("已购买", "warning");
    return;
  }
  if (userInfo.value.coins < item.coins) {
    addMessage("金币不足", "warning");
    return;
  }
  try {
    let { data, message } = await buyAttachment(item.id);
    init();
    addMessage(message, "success");
  } catch (error) {
    addMessage(error, "error");
  }
}
init();
</script>
