<template>
  <Dialog v-model="loginStatus" :confirmBtn="false">
    <form class="w-[400px] m-auto" @submit="handleSubmit">
      <!-- 密码登录 验证码登录 -->
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

      <!-- 账号 -->
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
        <!-- 发送验证码 -->
        <button
          type="button"
          v-if="loginType"
          @click="!countState ? (showCode = true) : () => {}"
          class="flex-shrink-0 text-xs text-light-5 cursor-pointer"
        >
          {{ countState ? count : $t("login-sendCode") }}
        </button>
      </div>
      <!-- 验证码 -->
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
          class="cursor-pointer text-[--font-secondary-color]"
          :name="pwdOptions[pwdType].name"
          size="1.5rem"
          @click="togglePasswordVisibility"
        ></Icon>
      </div>
      <div class="mt-4 grid grid-cols-2 gap-x-2">
        <button
          type="button"
          class="p-2 rounded-md col-span-1 shadow-center"
          v-if="!loginType"
          @click="changeLoginType(1)"
        >
          {{ $t("login-register") }}
        </button>
        <!-- 登录按钮 -->
        <button
          v-if="!loginType"
          type="button"
          @click="showCode = true"
          class="shadow-center p-2 w-full rounded-md text-white !bg-[--illuminate-color]"
        >
          {{ $t("login-login") }}
        </button>
        <button
          v-else
          type="button"
          @click="handleSubmit"
          class="shadow-center p-2 w-full rounded-md text-white !bg-[--illuminate-color] col-span-2"
        >
          {{ $t("login-login") }}
        </button>
      </div>
    </form>
  </Dialog>
  <Dialog v-model="showCode" :title="'验证码'" @confirm="startSendCode">
    <div class="flex rounded-lg shadow-center overflow-hidden m-5">
      <Captcha ref="captchaRef" @captchaId="setCaptchaId" class="h-10" />
      <input
        v-model="form.code"
        maxlength="4"
        class="w-20 p-2 focus:outline-none"
        type="text"
      />
    </div>
  </Dialog>
</template>

<script setup>
// 导入所需的API和状态管理函数
import { login, emailCode } from "~/api";
import { useLoginStore, useUserStore } from "~/stores/main.js";
const { notice } = useNotice();
const { setCookie } = useCookies();
const { sleep } = useSleep();

// 使用状态管理库
const loginStore = useLoginStore();
const userStore = useUserStore();

// 定义响应式变量
const loginStatus = computed({
  get: () => loginStore.getLoginStatus,
  set: (val) => loginStore.setLoginStatus(val),
});
let loading = ref(false);
let showCode = ref(false);
let captchaId = ref("");
let captchaRef = ref(null);

// 定义表单数据
let form = ref({
  username: "",
  password: "",
  code: "",
  vCode: "",
});

let loginType = ref(0); // 登录类型标记，0代表密码登录，1代表验证码登录
let pwdType = ref(0); // 密码可见性标记，0为隐藏，1为显示

// 密码输入框的配置选项
let pwdOptions = {
  0: {
    name: "ph:eye-closed-bold",
    type: "password",
    form: "password",
  },
  1: {
    name: "ph:eye-bold",
    type: "text",
    form: "vCode",
  },
};

let countState = ref(false); // 计时器状态标记
let count = ref(60); // 倒计时秒数
let timer = null; // 计时器变量

// 验证邮箱格式和空值
function validateEmail() {
  if (!form.value.username.includes("@") || form.value.username === "") {
    notice({
      title: "提示",
      content: "请输入正确的邮箱",
      autoClose: true,
    });
    return false;
  }
  return true;
}

// 组件销毁前清除计时器
onBeforeUnmount(() => {
  clearInterval(timer);
});

// 切换密码可见性
const togglePasswordVisibility = () => {
  pwdType.value = pwdType.value === 0 ? 1 : 0;
};

// 更改登录类型
const changeLoginType = (type) => {
  loginType.value = type;
  pwdType.value = type; // 重置密码可见性
};

// 设置验证码ID
const setCaptchaId = (id) => {
  captchaId.value = id;
};

// 处理表单提交
const handleSubmit = async () => {
  // if (loginType && !validateEmail()) return;

  // 确保用户名和密码/验证码不为空
  if (
    loginType.value === 0 &&
    (form.value.username === "" || form.value.password === "")
  ) {
    notice({
      title: "提示",
      content: "请输入用户名和密码",
      autoClose: true,
    });
    return;
  }

  // 构建请求对象
  let obj =
    loginType.value === 0
      ? { ...form.value, loginType: 1, captchaId: captchaId.value }
      : {
          email: form.value.username,
          code: form.value.code,
          vCode: form.value.vCode,
          loginType: 2,
          captchaId: captchaId.value,
        };

  try {
    // 发送登录请求
    let { data } = await login(obj);
    setCookie("token", data.token); // 设置token
    await sleep(500); // 稍作延迟
    await loginStore.setLoginStatus(); // 更新登录状态
    await userStore.fetchUserInfo(); // 获取用户信息

    // 显示登录成功通知
    notice({
      title: "提示",
      content: "登录成功",
      autoClose: true,
    });
  } catch (error) {
    notice({
      title: "提示",
      content: error,
      autoClose: true,
    });
    captchaRef.value.init(); // 重新初始化验证码
  }
};

// 发送验证码
const sendCode = async () => {
  if (!validateEmail()) return;

  try {
    // 请求发送验证码
    await emailCode({
      email: form.value.username,
      code: form.value.code,
      vCode: form.value.vCode,
      captchaId: captchaId.value,
    });

    // 启动计时器
    countState.value = true;
    clearInterval(timer); // 清除之前的计时器
    timer = setInterval(() => {
      if (count.value > 0) {
        count.value--;
      } else {
        countState.value = false;
        count.value = 60;
        clearInterval(timer);
      }
    }, 1000);
  } catch (error) {
    // 处理错误
    console.error(error);
  }
};

// 开始发送验证码的流程
const startSendCode = () => {
  loginType.value ? sendCode() : handleSubmit();
  form.value.code = ""; // 清空验证码输入框
};
</script>
