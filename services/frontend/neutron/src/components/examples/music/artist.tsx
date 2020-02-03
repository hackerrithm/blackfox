import React from "React";
import Artists from "./store";

const Artist: React.FC = () => {
	return (
		<>
			<h1>MTV Base Headline Artists 2019</h1>
			{Artists.map((item: any, index: number) => (
				<div id="card-body" key={index}>
					<div className="card">
						<h2>{item.name}</h2>
						<p>genre: {item.genre}</p>
						<p>Albums released: {item.albums}</p>
					</div>
				</div>
			))}
		</>
	);
};

export default Artist;
