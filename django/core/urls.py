from django.urls import path

from core.api import videos_list, videos_detail_by_id, videos_detail_by_slug

urlpatterns = [
    path('api/videos', videos_list, name='api_videos_list'),
    path('api/videos/<int:id>', videos_detail_by_id, name='api_videos_detail_by_id'),
    path('api/videos/<slug:slug>', videos_detail_by_slug, name='api_videos_detail_by_slug'),
]
