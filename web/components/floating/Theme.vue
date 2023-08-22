<template>
  <Floating :index="index" @change="colorPicker = !colorPicker">
    <div class="flex justify-between">
      <div class="bg-slate-400 rounded-2xl overflow-hidden w-1/2 text-center">
        Light
      </div>
      <div class="rounded-2xl w-1/2 text-center">Dark</div>
    </div>
    <div>
      <label>透明度</label>
      <input
        type="range"
        :value="getThemeValue.blurValue"
        @input="updateValue($event, 'blur-value')"
        min="0"
        max="20"
      />
    </div>
    <input type="color" />
    <div>
      <label>背景颜色</label>
      <ColorPicker
        v-if="colorPicker"
        :color="getThemeValue.bgColor"
        @change="changeColor"
      />
    </div>
    <Popover>
      <button>按钮</button>
      <template #content>
        <div>内容</div>
      </template>
    </Popover>
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
function changeColor(rgbaString) {
  getThemeValue.value.bgColor = rgbaString;
  changeCssVariables({
    "bg-color": rgbaString,
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
