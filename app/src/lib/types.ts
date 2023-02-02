export class CveSummary {
	id: string;
	description: string;
	published: string;
	cvssBaseScore: number;
	cvssBaseSeverity: string;
	cvssBaseVector: string;
	cvssSource: string;

	constructor(
		id: string,
		description: string,
		published: string,
		cvssBaseSeverity: string,
		cvssBaseScore: number,
		cvssBaseVector: string,
		cvssSource: string
	) {
		this.id = id;
		this.description = description;
		this.published = published;
		this.cvssBaseSeverity = cvssBaseSeverity;
		this.cvssBaseScore = cvssBaseScore;
		this.cvssBaseVector = cvssBaseVector;
		this.cvssSource = cvssSource;
	}
}
