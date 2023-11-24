<template>
  <div class="w-full flex justify-center my-5">
    <el-avatar :size="50" :src="''" />
  </div>
  <el-menu
    :default-active="getActive"
    class="w-[200px]"
    :collapse="props.modelValue"
    active-text-color="#ffd04b"
    background-color="#545c64"
    text-color="#fff"
    @select="select"
  >
    <component
      :is="item.children ? ElSubMenu : ElMenuItem"
      v-for="(item, index) in getRouter"
      :index="item.name"
    >
      <el-icon v-if="item.meta?.icon">
        <component :is="item.meta.icon"></component>
      </el-icon>
      <template #title>
        <el-icon v-if="item.meta?.icon && item.children">
          <component :is="item.meta.icon"></component>
        </el-icon>
        <span>{{ item.label }}</span>
      </template>
      <template #default v-if="item.children">
        <el-menu-item-group v-for="(val, i) in item.children">
          <el-menu-item :index="val.name">
            <template #title>
              <el-icon>
                <component
                  v-if="val.meta?.icon"
                  :is="val.meta.icon"
                ></component>
              </el-icon>
              {{ val.label }}
            </template>
          </el-menu-item>
        </el-menu-item-group>
      </template>
    </component>
  </el-menu>
</template>

<script setup>
import { computed, ref } from "vue";
import { useRouter } from "vue-router";
import { ElMenuItem, ElSubMenu } from "element-plus";
const router = useRouter();
const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false,
  },
});

const getRouter = computed(() => {
  // 辅助函数，用于决定是否包含路由
  const shouldIncludeRoute = (route) => {
    // 如果路由没有 meta 或 meta.show 为 true 或 undefined，返回 true
    return !route.meta || route.meta.show !== false;
  };

  // 递归函数处理路由及其子路由
  const processRoute = (route) => {
    // 判断当前路由是否应该包含
    if (shouldIncludeRoute(route)) {
      return {
        path: route.path,
        name: route.name,
        meta: route.meta,
        label: route.label,
        // 如果有子路由，递归处理并过滤
        children: route.children
          ? route.children.map(processRoute).filter((child) => child !== null)
          : undefined,
      };
    }
    return null;
  };

  // 遍历路由，应用处理函数，并过滤掉所有 null 值
  return router.options.routes[0].children
    .map(processRoute)
    .filter((route) => route !== null);
});
const getActive = computed(() => {
  return router.currentRoute.value.name;
});
const select = (key, keyPath) => {
  router.push({
    name: key,
  });
};
</script>

<style scoped>
.el-menu {
  border: none;
}
</style>
