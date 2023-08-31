import { Feed, FeedApiResponse } from "../components/Feed/interfaces";
import { instance } from "../services/api";
import { globalConstants, identityConstants } from "../utils/constants";

export async function create(feed: Feed): Promise<void> {
  var formData = new FormData();
  for (var key in feed) {
      formData.append(key, (feed as any)[key]);
  }

  try {
    if (
      !feed.name || !feed.url
    ) {
      throw new Error(identityConstants.FILL_REQUIRED_FIELDS);
    }

    await instance.post(`${globalConstants.BASE_URL}/feeds`, formData);
  } catch (error: any) {
    throw new Error(error.message);
  }
}

export async function getNotFollowedFeeds(): Promise<FeedApiResponse[]> {
  try {
    const result = await instance.get<FeedApiResponse[]>(`${globalConstants.BASE_URL}/feeds`);
    return result.data
  } catch (error: any) {
    throw new Error(error.message);
  }
}
