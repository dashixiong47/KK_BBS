// plugins/vconsole.client.js

import VConsole from 'vconsole';

export default () => {
  // 仅在开发模式下启用 vConsole
  if (process.env.NODE_ENV === 'development') {
    new VConsole();
  }
};
