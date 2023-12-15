<template>
  <video ref="video" controls></video>
</template>

<script>
import Hls from "hls.js";

export default {
  name: "VideoPlayer",
  props: {
    src: {
      type: String,
      required: true,
    },
  },
  mounted() {
    this.setupVideo();
  },
  watch: {
    src(newValue) {
      if (newValue) {
        this.setupVideo();
      }
    },
  },
  methods: {
    setupVideo() {
      const video = this.$refs.video;

      if (!this.src) {
        video.pause();
        video.src = ''; // 清除视频源
        return;
      }

      if (Hls.isSupported()) {
        const hls = new Hls();
        hls.loadSource(this.src);
        hls.attachMedia(video);
        hls.on(Hls.Events.MANIFEST_PARSED, () => {
          video.play();
        });
      } else if (video.canPlayType("application/vnd.apple.mpegurl")) {
        video.src = this.src;
        video.addEventListener("loadedmetadata", function () {
          this.play();
        });
      }
    }
  }
};
</script>
