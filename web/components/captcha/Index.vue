<template>
  <img :src="captchaUrl" alt="" srcset="" @click="init" />
</template>

<script setup>
import { captcha } from "~/api";
let captchaUrl = ref("");
let captchaId = ref("");
let emit = defineEmits(["captchaId"]);
async function init() {
  let { data } = await captcha();
  captchaUrl.value = data.captcha;
  captchaId.value = data.captchaId;
  emit("captchaId", captchaId.value);
}
init();
defineExpose({
  init,
});
</script>
