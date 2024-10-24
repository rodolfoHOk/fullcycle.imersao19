export type VideoModel = {
  id: number;
  title: string;
  description: string;
  thumbnail: string;
  slug: string;
  published_at: string;
  likes: number;
  views: number;
  tags: string[];
  video_url: string;
  // author: {
  //     id: number;
  //     name: string;
  //     avatar: string;
  // }
};
