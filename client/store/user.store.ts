import { defineStore } from 'pinia';
import type { User, UserRegister, UserLogin } from '~/types';
import { useApi } from '@/composables/useApi';

const api = useApi<User>('/user');

export const useUserStore = defineStore('user', {
  state: () => ({
    isLoggedIn: false,
    profile: {} as User
  }),
  actions: {
    async register(userData: UserRegister) {
      if (userData.password !== userData.confirmPassword) {
        throw new Error('Passwords do not match');
      }
      await api.create<UserRegister, User>(userData, '/register');
      this.isLoggedIn = true;
      this.persistLogin();
    },
    async login(userData: UserLogin) {
      const user: User = await api.create<UserLogin, User>(userData, '/login');
      this.isLoggedIn = true;
      this.profile = user;
    },
    async logout() {
      await api.create({}, '/logout');
      this.isLoggedIn = false;
      this.profile = {} as User;
    },
    persistLogin() {
      localStorage.setItem('auth', JSON.stringify({ isLoggedIn: this.isLoggedIn, profile: this.profile }));
    },
    async checkAuth() {
      try {
        const user = await api.get('/profile');
        this.isLoggedIn = true;
        this.profile = user;
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
      } catch (error) {
        this.isLoggedIn = false;
        this.profile = {} as User;
      }
    }
  }
});
