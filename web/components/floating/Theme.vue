<template>
  <Floating :index="index" @change="colorPicker = !colorPicker">
    <button
      class="relative z-10 bg-blue-500 hover:bg-blue-700 text-white font-bold w-10 h-10 rounded-full"
    >
      <Icon name="ion:color-palette-outline" size="1.5rem" />
    </button>
    <template #content>
      <div class=" overflow-y-auto h-full hidden-scrollbar">
        <div class="flex items-center justify-between">
          <label class="font-main-color w-20">透明度</label>
          <input
            type="range"
            :value="getThemeValue.blurValue"
            @input="updateValue($event, 'blur-value')"
            min="0"
            max="20"
          />
        </div>
        <div
          class="flex items-center justify-between"
          v-for="(val, key) of colorObj"
        >
          <label class="font-main-color w-32">{{ val }}</label>
          <Popover>
            <div
              class="w-10 h-5 border border-base"
              :style="{
                backgroundColor: getThemeValue[key],
              }"
            ></div>
            <template #content>
              <ColorPicker
                :color="getThemeValue[key]"
                @change="(v) => changeColor(key, v)"
              />
            </template>
          </Popover>
        </div>
      </div>
    </template>
  </Floating>
</template>

<script setup>
import { useThemeStore } from "~/stores/theme.js";
import useToKebabCase from "~/composables/useToKebabCase";
let { toKebabCase } = useToKebabCase();
const colorMode = useColorMode();
let themeType = computed(() => colorMode.value);
let themeStore = useThemeStore();
let unit = computed(() => themeStore.unit);
let config = computed(() => themeStore.getThemeConfig);
let darkConfig = computed(() => themeStore.getDarkThemeConfig);
let getThemeValue = computed({
  get() {
    if (themeType.value === "light") {
      return config.value;
    } else {
      return darkConfig.value;
    }
  },
  set(newV) {
    if (themeType.value === "light") {
      themeStore.setThemeConfig(newV);
    } else {
      themeStore.setDarkThemeConfig(newV);
    }
  },
});
let colorPicker = ref(false);
let { index } = defineProps({
  index: {
    type: Number,
    default: 0,
  },
});
let colorObj = {
  bgColor: "背景颜色", // 背景颜色
  // 默认主字体颜色
  fontMainColor: "默认主字体颜色",
  // 常规文本颜色
  fontRegularColor: "常规文本颜色",
  // 次要文本颜色
  fontSecondaryColor: "次要文本颜色",
  // 选中背景色
  activeBgColor: "选中背景色",
  // 选中字体颜色
  activeTextColor: "选中字体颜色",
  // 未选中背景色
  inactiveBgColor: "未选中背景色",
  // 未选中字体颜色
  inactiveTextColor: "未选中字体颜色",
  // 禁用色
  bgDisableColor: "禁用色",
  //禁用背景色
  fontDisableBgColor: "禁用背景色",
  
};
useHead({
  script: {
    class: "irojs",
    src: "/iro.js",
  },
  style: [
    {
      type: "text/css",
      class: "light",
      innerHTML: createCssVariables(getCssVariables(config.value)),
    },
    {
      type: "text/css",
      class: "dark",
      innerHTML: createCssVariables(getCssVariables(darkConfig.value), "dark-"),
    },
  ],
});

function updateValue(event, key) {
  // 判断event是否是对象
  let value = event.target.value;
  switch (key) {
    case "blur-value":
      getThemeValue.value.blurValue = value;
      value = value + "px";

      break;
    case "bg-color":
      getThemeValue.value.bgColor = value;
      break;

    default:
      break;
  }
  changeCssVariables({
    [key]: value,
  });
}
function changeColor(name, rgbaString) {
  getThemeValue.value[name] = rgbaString;
  changeCssVariables({
    [toKebabCase(name)]: rgbaString,
  });
}
/*
 * 获取 css 变量
 * @param {Object} css
 * @returns {Object}
 */
function getCssVariables(css) {
  let cssVariables = {};
  for (const key in css) {
    cssVariables[toKebabCase(key)] = unit.value[key]
      ? css[key] + unit.value[key]
      : css[key];
  }
  return cssVariables;
}
/**
 * 修改 css 变量
 * @param {Object} { key: value}
 * @description { key: value} = { --blur-value: 10px }
 */
function changeCssVariables(variables) {
  requestAnimationFrame(() => {
    const prefix = themeType.value === "light" ? "" : `${themeType.value}-`;
    let styleChanges = "";
    for (const [key, value] of Object.entries(variables)) {
      styleChanges += `--${prefix}${key}: ${value}; `;
    }
    document.documentElement.style.cssText += styleChanges;
  });
}

/**
 * 把 css 变量转换成字符串
 * @param {Object} { key: value}
 * @returns {string}
 * @description { key: value} = { --blur-value: 10px }
 */
function createCssVariables(variables, themeType = "") {
  let cssString = ":root {";

  for (const [key, value] of Object.entries(variables)) {
    cssString += `--${themeType}${key}: ${value};`;
  }

  cssString += "}";
  return cssString;
}
</script>
