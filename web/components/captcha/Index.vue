<template>
  <img :src="captchaUrl" alt="" srcset="" />
</template>

<script setup>
let { captcha } = useApi();
let captchaUrl = ref("");
let captchaId = ref("");
async function init() {
  let res = await captcha();
  console.log(res);
  captchaUrl.value = res.captcha;
}
init();

</script>
