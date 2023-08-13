<template>
  <nav
    class="shadow-md border-b h-16 box-content flex-shrink-0 content-center px-5 grid grid-cols-2 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4"
  >
    <div class="md:col-span-1">
      <KLink to="/" class="flex items-center h-full">
        <img
          class="h-8 mr-5"
          src="https://hxsj.in/_nuxt/img/logo.5c704a9.png"
          alt=""
          srcset=""
        />
        KK_BBS
      </KLink>
    </div>
    <div class="hidden lg:flex items-center col-span-2">
      <Search />
      <ul class="w-4/6 h-full flex items-center px-5 text-blue-400">
        <li class="mr-2">
          <button class="bg-gray-300 rounded-full w-10 h-10">
            <Icon name="icon-zhuye" size="text-2xl mx-1"></Icon>
          </button>
        </li>
        <li class="mr-2">
          <button class="bg-gray-300 rounded-full w-10 h-10">
            <Icon name="icon-neirong" size="text-2xl mx-1"></Icon>
          </button>
        </li>
        <li class="mr-2">
          <button class="bg-gray-300 rounded-full w-10 h-10">
            <Icon name="icon-video" size="text-2xl mx-1"></Icon>
          </button>
        </li>
        <li class="mr-2">
          <button class="bg-gray-300 rounded-full w-10 h-10">
            <Icon name="icon-wendang" size="text-2xl mx-1"></Icon>
          </button>
        </li>
      </ul>
    </div>
    <div class="md:col-span-1 flex items-center justify-end">
      <Icon name="icon-pinglun" size="text-2xl mx-2"></Icon>
      <Icon name="icon-xiaoxi" size="text-2xl mx-2"></Icon>
      <Icon name="icon-shezhi" size="text-2xl mx-2"></Icon>
      <Switcher />
      <KLink :to="`/user/1`">
        <Avatar class="ml-4" />
      </KLink>
      <!-- <div class="absolute">1212</div> -->
    </div>
  </nav>
</template>

<script setup></script>
