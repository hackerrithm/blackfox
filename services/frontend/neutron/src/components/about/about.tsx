import React from "react";
import { Suspense, lazy } from "react";
const Artist = React.lazy(() => import("../music/artist"));
const Performers = lazy(() => import("../music/performers"));

const About = () => {
	return (
		<div>
			<h2>About</h2>
			<Suspense fallback={<div>Loading...</div>}>
				<Performers />
				<Artist />
			</Suspense>
		</div>
	);
};

export default About;
