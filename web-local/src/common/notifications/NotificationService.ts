import type { NotificationOptions } from "@rilldata/web-local/lib/components/notifications/notificationStore";

export interface Notification {
  message: string;
  type: string;
  detail?: string;
  options?: NotificationOptions;
}

export interface NotificationService {
  notify(notification: Notification): void;
}
