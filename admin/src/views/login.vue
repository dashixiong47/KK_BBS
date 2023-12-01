<template>
  <div class="w-full h-[100vh] flex justify-center">
    <div class="w-[500px] h-96 mt-36 p-5 border rounded-xl">
      <h2 class="text-center mb-5">登录</h2>
      <el-form :model="form">
        <el-form-item>
          <el-input
            v-model="form.username"
            placeholder="请输入账号"
            clearable
          />
        </el-form-item>
        <el-form-item>
          <el-input
            v-model="form.password"
            placeholder="请输入密码"
            type="password"
            clearable
          />
        </el-form-item>
        <el-form-item>
          <div class="flex items-center">
            <el-input
              v-model="form.code"
              placeholder="请输入验证码"
              clearable
            />
            <img
              class="h-8 ml-5"
              @click="init"
              :src="codeBase64"
              alt=""
              srcset=""
            />
          </div>
        </el-form-item>
        <el-button type="primary" @click="loginEvt" style="width: 100%">
          登录
        </el-button>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { captcha, login } from "@/api/user";
// import { useI18n } from "vue-i18n";
// const { t } = useI18n();
let router = useRouter();
let form = ref({
  username: "",
  password: "",
  code: "",
  captchaId: "",
});
let codeBase64 = ref("");
async function init() {
  try {
    let { data } = await captcha();
    form.value.captchaId = data.captchaId;
    codeBase64.value = data.captcha;
  } catch (error) {
    console.log(error);
  }
}
async function loginEvt() {
  try {
    let { data } = await login(form.value);
    localStorage.setItem("token", data.token);
    router.replace({ path: "/" });
  } catch (error) {
    console.log(error);
  }
}
init();
</script>
