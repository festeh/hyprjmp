import { writable } from 'svelte/store';

export type Mode = 'jump' | 'book';

export const modeStore = writable<Mode>('jump');
