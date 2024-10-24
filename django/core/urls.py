from django.urls import path

from core.api import videos_list

urlpatterns = [
    path('api/videos', videos_list, name='api_videos_list'),
]
