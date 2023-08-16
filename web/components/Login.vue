<template>
  <div class="bg-[rgba(0,0,0,0.6)] fixed top-0 left-0 w-screen h-screen z-50">
    <div
      class="p-5 bg-white rounded-2xl absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2"
    >
      <p class="w-full flex justify-end">
        <Icon class="cursor-pointer" name="icon-close" @click="close" />
      </p>

      <form class="w-[400px] m-auto">
        <div class="mt-4 mb-4 flex justify-center">
          <div
            class="cursor-pointer text-center text-lg pr-5 border-r"
            :class="{
              'text-blue-400': !loginType,
            }"
            @click="changeLoginType(0)"
          >
            密码登录
          </div>
          <div
            class="cursor-pointer text-center text-lg pl-5"
            :class="{
              'text-blue-400': loginType,
            }"
            @click="changeLoginType(1)"
          >
            验证码登录
          </div>
        </div>
        <div
          class="h-11 p-2 border rounded-t-md text-sm overflow-hidden flex items-center"
        >
          <label class="flex-shrink-0 mr-2 w-12 text-center">账号</label>
          <input
            type="text"
            class="w-full focus:outline-none"
            placeholder="请输入邮箱账号"
          />
          <button v-if="loginType" class="flex-shrink-0 text-xs text-light-5 cursor-pointer">发送验证码</button>
        </div>
        <div
          class="h-11 p-2 border rounded-b-md border-t-0 text-sm overflow-hidden flex items-center"
        >
          <label class="flex-shrink-0 mr-2 w-12 text-center">
            {{ loginType ? "验证码" : "密码" }}
          </label>
          <input
            :type="pwdOptions[pwdType].type"
            class="w-full focus:outline-none"
            :placeholder="!loginType ? '请输入密码' : '请输入验证码'"
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
            class="border p-2 rounded-md col-span-1"
            v-if="!loginType"
            @click="changeLoginType(1)"
          >
            注册
          </button>
          <button
            class="border p-2 rounded-md col-span-1 text-white bg-blue-400"
            :class="{
              'col-span-2': loginType,
            }"
          >
            登录
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { useLoginStore } from "~/stores/main.js";
const store = useLoginStore();

let loginType = ref(0);
let pwdType = ref(0);
let pwdOptions = {
  0: {
    name: "icon-bukejian",
    type: "password",
  },
  1: {
    name: "icon-kejian",
    type: "text",
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
</script>
