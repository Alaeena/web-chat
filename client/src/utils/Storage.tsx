import { UserInfo } from '@/modules/auth_provider';

const KEY = 'CHAT_APP_STORAGE';
const initState = {
  id: '',
  username: '',
};

// eslint-disable-next-line import/no-anonymous-default-export
export default {
  get(): UserInfo {
    var value = localStorage.getItem(KEY);
    return (value && JSON.parse(value)) || initState;
  },
  set(value: UserInfo) {
    localStorage.setItem(KEY, JSON.stringify(value));
  },
};
