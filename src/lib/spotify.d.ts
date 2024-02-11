export interface Track {
	id: string;
	BPM: number;
	key: number;
	url: string;
	timeSignature: number;
	previewURL: string;
	image: string;
	name: string;
	artistName: string;
}

export interface Playlist {
	tracks: Track[];
}
