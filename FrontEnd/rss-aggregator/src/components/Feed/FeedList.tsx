import { useEffect, useState } from "react";
import { FeedApiResponse } from "./interfaces";
  
import * as feedService from "../../services/feedService";

interface Props {
    followed: boolean,
}

const FeedList: React.FC<Props> = (params) => {
    const [feeds, setFeeds] = useState<FeedApiResponse[]>([]);

    const { followed } = params;

    useEffect(() => {
        if (followed) {

        } else {
            feedService.getNotFollowedFeeds()
            .then(res => {
                setFeeds(res);
            });
        }
    }, [followed])

  return (
    <div className="cars-list-wrapper">
      {feeds.map((feed) => {
        return (
          <FeedCard
            key={feed.id}
            feed={feed}
            followFeed={followed ? followFeed : undefined}
            unfollowFeed={followed ? undefined : unfollowFeed}
          />
        );
      })}
    </div>
  );
};

export default FeedList;