<template>
  <div
    v-if="loginStatus"
    class="bg-[rgba(0,0,0,0.6)] fixed top-0 left-0 w-screen h-screen z-50"
  >
    <div
      class="p-5 bg-white rounded-2xl absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2"
    >
      <p class="w-full flex justify-end">
        <Icon class="cursor-pointer" name="icon-close" @click="close" />
      </p>

      <form class="w-[400px] m-auto" @submit="handleSubmit">
        <div class="mt-4 mb-4 flex justify-center">
          <div
            class="cursor-pointer text-center text-lg pr-5 border-r"
            :class="{
              'text-blue-400': !loginType,
            }"
            @click="changeLoginType(0)"
          >
            {{ $t("login-passwordLogin") }}
          </div>
          <div
            class="cursor-pointer text-center text-lg pl-5"
            :class="{
              'text-blue-400': loginType,
            }"
            @click="changeLoginType(1)"
          >
            {{ $t("login-codeLogin") }}
          </div>
        </div>
        <div
          class="h-11 p-2 border rounded-t-md text-sm overflow-hidden flex items-center"
        >
          <label class="flex-shrink-0 mr-2 w-12 text-center">{{
            $t("login-account")
          }}</label>
          <input
            type="text"
            v-model="form.username"
            class="w-full focus:outline-none"
            :placeholder="$t('login-enterYourAccount')"
          />
          <button
            v-if="loginType"
            class="flex-shrink-0 text-xs text-light-5 cursor-pointer"
          >
            {{ $t("login-sendCode") }}
          </button>
        </div>
        <div
          class="h-11 p-2 border rounded-b-md border-t-0 text-sm overflow-hidden flex items-center"
        >
          <label class="flex-shrink-0 mr-2 w-12 text-center">
            {{ loginType ? $t("login-code") : $t("login-password") }}
          </label>
          <input
            v-model="form[pwdOptions[pwdType].form]"
            :type="pwdOptions[pwdType].type"
            class="w-full focus:outline-none"
            :placeholder="
              !loginType
                ? $t('login-enterYourPassword')
                : $t('login-enterYourCode')
            "
          />
          <Icon
            v-if="!loginType"
            class="text-light-5 cursor-pointer"
            :name="pwdOptions[pwdType].name"
            @click="show"
          ></Icon>
        </div>
        <div class="mt-4 grid grid-cols-2 gap-x-2">
          <button
            type="button"
            class="border p-2 rounded-md col-span-1"
            v-if="!loginType"
            @click="changeLoginType(1)"
          >
            {{ $t("login-register") }}
          </button>
          <button
            class="border p-2 rounded-md col-span-1 text-white bg-blue-400"
            :class="{
              'col-span-2': loginType,
            }"
          >
            {{ $t("login-login") }}
          </button>
        </div>
        <Captcha />
      </form>
    </div>
  </div>
</template>

<script setup>
import { useLoginStore } from "~/stores/main.js";
const loginStatus = computed(() => store.getLoginStatus);
let { login } = useApi();
const store = useLoginStore();
const { notice } = useNotice();

let form = ref({
  username: "",
  password: "",
  code: "",
});
let loginType = ref(0);
let pwdType = ref(0);
let pwdOptions = {
  0: {
    name: "icon-bukejian",
    type: "password",
    form: "password",
  },
  1: {
    name: "icon-kejian",
    type: "text",
    form: "code",
  },
};

const close = () => {
  store.setLoginStatus();
};
const show = () => {
  pwdType.value = pwdType.value === 0 ? 1 : 0;
};
const changeLoginType = (type) => {
  if (type) {
    pwdType.value = 1;
  } else {
    pwdType.value = 0;
  }
  loginType.value = type;
};
const handleSubmit = async (event) => {
  event.preventDefault(); // 阻止表单的默认提交行为
  switch (loginType.value) {
    case 0:
      break;

    default:
      break;
  }
  try {
    let data = await login(form.value);
    localStorage.setItem("token", data.token);
    store.setLoginStatus();
    notice({
      title: "提示",
      content: "登录成功",
      autoClose: true,
    });
  } catch (error) {
    notice({
      title: "提示",
      content: "用户名或密码错误",
      autoClose: true,
    });
  }
};
</script>
