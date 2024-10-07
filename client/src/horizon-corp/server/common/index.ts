// index.ts
import ContactService from './ContactService';
import GenderService from './GenderService';
import RolesService from './RolesService';
import FeedbackService from './FeedbackService';
import Connection from './Connection';

export default function common() {
  return {
    contact: new ContactService(),
    gender: new GenderService(),
    roles: new RolesService(),
    feedback: new FeedbackService(),
    connection: new Connection(),
  };
}
