import { writable } from 'svelte/store';

export const tableview = writable<boolean>(
	typeof window !== 'undefined' && localStorage.tableview === 'true'
);

tableview.subscribe((value) => {
	if (typeof window !== 'undefined') localStorage.tableview = String(value);
});
