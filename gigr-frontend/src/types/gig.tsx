export interface Gig {
    id: number;
    userId: number;
    title: string;
    description: string;
    location: string;
    date: string;
    instrument: string;
    status: string;
    user?: {
        id: number;
        name: string;
        email: string;
  };
}