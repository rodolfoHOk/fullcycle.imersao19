from django.urls import path

from core.api import videos_list, videos_detail_by_id, videos_detail_by_slug, videos_list_recommended, videos_get_likes, videos_get_views

urlpatterns = [
    path('api/videos', videos_list, name='api_videos_list'),
    path('api/videos/<int:id>', videos_detail_by_id, name='api_videos_detail_by_id'),
    path('api/videos/<slug:slug>', videos_detail_by_slug, name='api_videos_detail_by_slug'),
    path('api/videos/<int:id>/recommended', videos_list_recommended, name='api_videos_list_recommended'),
    path('api/videos/<int:id>/likes', videos_get_likes, name='api_videos_get_likes'),
    path('api/videos/<int:id>/views', videos_get_views, name='api_videos_get_views'),
]
