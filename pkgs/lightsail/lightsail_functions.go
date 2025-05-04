package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
)

type Client struct {
	ls *lightsail.Client
}

func NewClient() (*Client, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	return &Client{
		ls: lightsail.NewFromConfig(cfg),
	}, nil
}

func (c *Client) StartInstance(ctx context.Context, instanceName string) error {
	_, err := c.ls.StartInstance(ctx, &lightsail.StartInstanceInput{
		InstanceName: &instanceName,
	})
	return err
}

func (c *Client) StopInstance(ctx context.Context, instanceName string) error {
	_, err := c.ls.StopInstance(ctx, &lightsail.StopInstanceInput{
		InstanceName: &instanceName,
	})
	return err
}

func (c *Client) RebootInstance(ctx context.Context, instanceName string) error {
	_, err := c.ls.RebootInstance(ctx, &lightsail.RebootInstanceInput{
		InstanceName: &instanceName,
	})
	return err
}

func (c *Client) GetInstanceState(ctx context.Context, instanceName string) (string, error) {
	resp, err := c.ls.GetInstance(ctx, &lightsail.GetInstanceInput{
		InstanceName: &instanceName,
	})
	if err != nil {
		return "", err
	}
	return string(*resp.Instance.State.Name), nil
}
