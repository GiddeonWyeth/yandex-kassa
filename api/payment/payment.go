package payment

import (
	"context"
	"encoding/json"

	"github.com/dbzer0/yandex-kassa/api/client"
)

type Payment struct {
	APIClient    *client.APIClient `json:"-"`
	ID           string            `json:"id"`                       // идентификатор платежа в Яндекс.Кассе
	Status       string            `json:"status"`                   // статус платежа. Возможные значения: pending, waiting_for_capture, succeeded и canceled
	Amount       Amount            `json:"amount"`                   // сумма платежа
	Description  *string           `json:"description,omitempty"`    // описание транзакции (не более 128 символов), которое вы увидите в личном кабинете Яндекс.Кассы
	Recipient    Recipient         `json:"recipient"`                // получатель платежа
	Requestor    Requestor         `json:"requestor"`                // инициатор платежа или возврата
	Method       *Method           `json:"payment_method,omitempty"` // способ оплаты, который был использован для платежа
	CreatedAt    string            `json:"created_at"`               // время создания заказа в формате ISO 8601. Пример: 2017-11-03T11:52:31.827Z
	Test         bool              `json:"test"`                     // признак тестовой операции
	Paid         bool              `json:"paid"`                     // признак оплаты заказа
	Refundable   bool              `json:"refundable"`               // возможность провести возврат по API
	Confirmation *Confirmation     `json:"confirmation,omitempty"`   // данные, необходимые для инициации выбранного сценария подтверждения платежа пользователем
}

type Amount struct {
	Value    string `json:"value"`    // сумма в выбранной валюте. Выражается в виде строки и пишется через точку
	Currency string `json:"currency"` // код валюты в формате ISO-4217
}

type Recipient struct {
	AccountID *string `json:"account_id,omitempty"` // идентификатор магазина в Яндекс.Кассе
	GatewayID *string `json:"gateway_id,omitempty"` // идентификатор субаккаунта. Используется для разделения потоков платежей в рамках одного аккаунта
}

// Payment создает создает объект Payment по которому доступны операции:
//   * получения информации о платеже;
//   * подтверждение платежа;
//   * отмена платежа;
func New(apiClient *client.APIClient, paymentID string) *Payment {
	return &Payment{
		ID:        paymentID,
		APIClient: apiClient,
	}
}

// Find позволяет получить информацию о текущем состоянии платежа по
// его уникальному идентификатору.
func (p *Payment) Find(ctx context.Context) (*Payment, error) {
	reply, err := p.APIClient.PaymentFind(ctx, p.ID, nil)
	if err != nil {
		return nil, err
	}
	defer reply.Close()

	if err := json.NewDecoder(reply).Decode(&p); err != nil {
		return nil, err
	}

	return p, nil
}

// Capture подтверждает вашу готовность принять платеж.
func (p *Payment) Capture() (*Payment, error) {
	return p, nil
}

// Cancel отменяет платеж, находящийся в статусе waiting_for_capture.
func (p *Payment) Cancel() (*Payment, error) {
	return p, nil
}
