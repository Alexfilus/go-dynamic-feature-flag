package test

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/rueidis"
)

type DynamicConfig struct {
	client        rueidis.Client
	testStr1      string
	testStr2      string
	testDuration1 time.Duration
	testDuration2 time.Duration
	testInt1      int
	testInt2      int
	testBool1     bool
	testBool2     bool
}

const (
	strKeyTestStr2           = "testStr2"
	strKeyTestStr1           = "testStr1"
	durationKeyTestDuration1 = "testDuration1"
	durationKeyTestDuration2 = "testDuration2"
	intKeyTestInt1           = "testInt1"
	intKeyTestInt2           = "testInt2"
	boolKeyTestBool1         = "testBool1"
	boolKeyTestBool2         = "testBool2"
)

func NewDynamicConfig(client rueidis.Client) *DynamicConfig {
	return &DynamicConfig{
		client:        client,
		testStr1:      "test1",
		testStr2:      "test2",
		testDuration1: 1000 * time.Millisecond,
		testDuration2: 120000 * time.Millisecond,
		testInt1:      1,
		testInt2:      2,
		testBool1:     true,
		testBool2:     false,
	}
}

func (c *DynamicConfig) TestStr1(ctx context.Context) string {
	resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(strKeyTestStr1).Cache(),
		time.Minute,
	).ToString()

	if err != nil {
		return c.testStr1
	}

	c.testStr1 = resp
	return resp
}

func (c *DynamicConfig) SetTestStr1(value string) *DynamicConfig {
	c.testStr1 = value
	return c
}

func (c *DynamicConfig) StoreTestStr1(ctx context.Context, value string) error {
	return c.client.Do(
		ctx,
		c.client.B().Set().Key(strKeyTestStr1).
			Value(value).
			Build()).
		Error()
}

func (c *DynamicConfig) TestStr2(ctx context.Context) string {
	resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(strKeyTestStr2).Cache(),
		time.Minute,
	).ToString()

	if err != nil {
		return c.testStr2
	}

	c.testStr2 = resp
	return resp
}

func (c *DynamicConfig) SetTestStr2(value string) *DynamicConfig {
	c.testStr2 = value
	return c
}

func (c *DynamicConfig) StoreTestStr2(ctx context.Context, value string) error {
	return c.client.Do(
		ctx,
		c.client.B().Set().Key(strKeyTestStr2).
			Value(value).
			Build()).
		Error()
}

func (c *DynamicConfig) TestDuration1(ctx context.Context) time.Duration {
	resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(durationKeyTestDuration1).Cache(),
		time.Minute,
	).AsInt64()

	if err != nil {
		return c.testDuration1
	}

	c.testDuration1 = time.Duration(resp) * time.Millisecond
	return c.testDuration1
}

func (c *DynamicConfig) SetTestDuration1(value time.Duration) *DynamicConfig {
	c.testDuration1 = value
	return c
}

func (c *DynamicConfig) StoreTestDuration1(ctx context.Context, value time.Duration) error {
	return c.client.Do(
		ctx,
		c.client.B().Set().Key(durationKeyTestDuration1).
			Value(strconv.FormatInt(value.Milliseconds(), 10)).
			Build()).
		Error()
}

func (c *DynamicConfig) TestDuration2(ctx context.Context) time.Duration {
	resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(durationKeyTestDuration2).Cache(),
		time.Minute,
	).AsInt64()

	if err != nil {
		return c.testDuration2
	}

	c.testDuration2 = time.Duration(resp) * time.Millisecond
	return c.testDuration2
}

func (c *DynamicConfig) SetTestDuration2(value time.Duration) *DynamicConfig {
	c.testDuration2 = value
	return c
}

func (c *DynamicConfig) StoreTestDuration2(ctx context.Context, value time.Duration) error {
	return c.client.Do(
		ctx,
		c.client.B().Set().Key(durationKeyTestDuration2).
			Value(strconv.FormatInt(value.Milliseconds(), 10)).
			Build()).
		Error()
}

func (c *DynamicConfig) TestInt1(ctx context.Context) int {
	resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(intKeyTestInt1).Cache(),
		time.Minute,
	).AsInt64()

	if err != nil {
		return c.testInt1
	}

	c.testInt1 = int(resp)
	return c.testInt1
}

func (c *DynamicConfig) SetTestInt1(value int) *DynamicConfig {
	c.testInt1 = value
	return c
}

func (c *DynamicConfig) StoreTestInt1(ctx context.Context, value int) error {
	return c.client.Do(
		ctx,
		c.client.B().Set().Key(intKeyTestInt1).
			Value(strconv.Itoa(value)).
			Build()).
		Error()
}

func (c *DynamicConfig) TestInt2(ctx context.Context) int {
	resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(intKeyTestInt2).Cache(),
		time.Minute,
	).AsInt64()

	if err != nil {
		return c.testInt2
	}

	c.testInt2 = int(resp)
	return c.testInt2
}

func (c *DynamicConfig) SetTestInt2(value int) *DynamicConfig {
	c.testInt2 = value
	return c
}

func (c *DynamicConfig) StoreTestInt2(ctx context.Context, value int) error {
	return c.client.Do(
		ctx,
		c.client.B().Set().Key(intKeyTestInt2).
			Value(strconv.Itoa(value)).
			Build()).
		Error()
}

func (c *DynamicConfig) TestBool1(ctx context.Context) bool {
	resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(boolKeyTestBool1).Cache(),
		time.Minute,
	).AsBool()

	if err != nil {
		return c.testBool1
	}

	c.testBool1 = resp
	return resp
}

func (c *DynamicConfig) SetTestBool1(value bool) *DynamicConfig {
	c.testBool1 = value
	return c
}

func (c *DynamicConfig) StoreTestBool1(ctx context.Context, value bool) error {
	return c.client.Do(
		ctx,
		c.client.B().Set().Key(boolKeyTestBool1).
			Value(strconv.FormatBool(value)).
			Build()).
		Error()
}

func (c *DynamicConfig) TestBool2(ctx context.Context) bool {
	resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(boolKeyTestBool2).Cache(),
		time.Minute,
	).AsBool()

	if err != nil {
		return c.testBool2
	}

	c.testBool2 = resp
	return resp
}

func (c *DynamicConfig) SetTestBool2(value bool) *DynamicConfig {
	c.testBool2 = value
	return c
}

func (c *DynamicConfig) StoreTestBool2(ctx context.Context, value bool) error {
	return c.client.Do(
		ctx,
		c.client.B().Set().Key(boolKeyTestBool2).
			Value(strconv.FormatBool(value)).
			Build()).
		Error()
}
