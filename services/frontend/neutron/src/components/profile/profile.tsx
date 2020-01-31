import React from "react";
import { Grid, Paper, makeStyles, Theme, createStyles } from "@material-ui/core";
import { getThemeProps } from "@material-ui/styles";
import MediaCard from "../general/reusable/card/simpleCard";
import BasicInfoCard from "../general/reusable/card/basicInfoCard";
import DevUpdateToggler from "./userInfoToggler";
import InfluenceMetricSummary from "./followersCounter";

const useStylesProfileSection = makeStyles((theme: Theme) =>
	createStyles({
		root: {
			flexGrow: 1
		},
		paper: {
			padding: theme.spacing(2),
			textAlign: "center",
			color: theme.palette.text.secondary
		},
		paperWithBackgroundIMG: {
			backgroundImage: "url(https://images.wallpaperscraft.com/image/code_letters_screen_137590_3840x2160.jpg)",
			padding: theme.spacing(2),
			textAlign: "center",
			color: theme.palette.text.secondary
		}
		
	})
);

const Profile = () => {
	const classes = useStylesProfileSection(getThemeProps);
	return (
		<div>
			<Grid container spacing={3}>
				<Grid item xs={12} sm={12}>
					<Paper className={classes.paper}>
						
					</Paper>
				</Grid>
				<Grid item xs={12} sm={6}>
					<Paper className={classes.paperWithBackgroundIMG}>
						<MediaCard imageURL={"https://hips.hearstapps.com/ell.h-cdn.co/assets/16/23/elle-hazel-eyes-mila.jpg"} />
					</Paper>
					<div className={"basic-info"}>
						<InfluenceMetricSummary followers={1200654} following={1} level={9} influence={4.7}  />
						<br/>
						<BasicInfoCard searchItemURL={"About me"} searchItemID={"I am great"}/>
						<br/>
						<BasicInfoCard searchItemURL={"My Stack"} searchItemID={"Golang, TypesScript"}/>
						<br/>
						<BasicInfoCard searchItemURL={"Hire me to"} searchItemID={"Do cool shit"}/>
					</div>
				</Grid>
				<Grid item xs={12} sm={6}>
					<DevUpdateToggler />
				</Grid>
				{/* {goalsList.map((val, index) => {
					return (
						<Grid
							alignContent={"center"}
							item
							xs={6}
							sm={3}
							key={index}
						>
							<ExpandableCard />
						</Grid>
					);
				})} */}
			</Grid>
		</div>
	);
};

export default Profile;
