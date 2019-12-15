import React, { useState, useEffect } from "react";
import { Suspense, lazy } from "react";
const Artist = React.lazy(() => import("../music/artist"));
import EmblaCarouselReact from "embla-carousel-react";
import Performers from "../music/performers";

const About = () => {
	const [embla, setEmbla] = useState(null);

	useEffect(() => {
		if (embla) {
			embla.on("select", () => {
				console.log(`Current index is ${embla.selectedScrollSnap()}`);
			});
		}
	}, [embla]);

	return (
		<div>
			<h2>About</h2>
			<EmblaCarouselReact emblaRef={setEmbla} options={{ loop: true }}>
				<div style={{ display: "flex" }}>
					<div style={{ flex: "0 0 100%" }}>
						<img
							alt={"1"}
							width="500px"
							height="200px"
							src={
								"https://66.media.tumblr.com/e7b8cf45f9d29ffe6644ebeaf9e87419/tumblr_oy6vcn7U7l1vtzqkfo1_400.jpg"
							}
						/>
					</div>
					<div style={{ flex: "0 0 100%" }}>
						{" "}
						<img
							alt={"2"}
							height="200px"
							width="500px"
							src={
								"https://i.ytimg.com/vi/pLqipJNItIo/hqdefault.jpg?sqp=-oaymwEYCNIBEHZIVfKriqkDCwgBFQAAiEIYAXAB&rs=AOn4CLBkklsyaw9FxDmMKapyBYCn9tbPNQ"
							}
						/>
					</div>
					<div style={{ flex: "0 0 100%" }}>
						{" "}
						<img
							alt={"3"}
							height="200px"
							width="500px"
							src={
								"https://i.ytimg.com/vi/kkLk2XWMBf8/hqdefault.jpg?sqp=-oaymwEYCNIBEHZIVfKriqkDCwgBFQAAiEIYAXAB&rs=AOn4CLB4GZTFu1Ju2EPPPXnhMZtFVvYBaw"
							}
						/>
					</div>
				</div>
			</EmblaCarouselReact>
			<button onClick={() => embla.scrollPrev()}>Prev</button>
			<button onClick={() => embla.scrollNext()}>Next</button>
			<Suspense fallback={<div>Loading...</div>}>
				<Performers />
				<Artist />
			</Suspense>
		</div>
	);
};

export default About;
