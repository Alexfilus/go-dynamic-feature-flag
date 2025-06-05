package test

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/rueidis"
)

type DynamicConfig struct {
	client rueidis.Client
	testStr1 string
	testStr2 string
	testDuration1 time.Duration
	testDuration2 time.Duration
	testInt1 int
	testInt2 int
	testBool1 bool
	testBool2 bool
}

type RequestDynamicConfigUpdate struct {
	TestStr1 *string `json:"test_str1,omitempty"`
	TestStr2 *string `json:"test_str2,omitempty"`
	TestDuration1 *time.Duration `json:"test_duration1,omitempty"`
	TestDuration2 *time.Duration `json:"test_duration2,omitempty"`
	TestInt1 *int `json:"test_int1,omitempty"`
	TestInt2 *int `json:"test_int2,omitempty"`
	TestBool1 *bool `json:"test_bool1,omitempty"`
	TestBool2 *bool `json:"test_bool2,omitempty"`
}

type ResponseDynamicConfig struct {
	TestStr1 string `json:"test_str1"`
	TestStr2 string `json:"test_str2"`
	TestDuration1 time.Duration `json:"test_duration1"`
	TestDuration2 time.Duration `json:"test_duration2"`
	TestInt1 int `json:"test_int1"`
	TestInt2 int `json:"test_int2"`
	TestBool1 bool `json:"test_bool1"`
	TestBool2 bool `json:"test_bool2"`
}

const (
	strKeyTestStr1 = "test-project:test:test_str1"
	strKeyTestStr2 = "test-project:test:test_str2"
	durationKeyTestDuration1 = "test-project:test:test_duration1"
	durationKeyTestDuration2 = "test-project:test:test_duration2"
	intKeyTestInt1 = "test-project:test:test_int1"
	intKeyTestInt2 = "test-project:test:test_int2"
	boolKeyTestBool2 = "test-project:test:test_bool2"
	boolKeyTestBool1 = "test-project:test:test_bool1"
)

func NewDynamicConfig(client rueidis.Client) *DynamicConfig {
	return &DynamicConfig{
		client: client,
		testStr1: "test1",
		testStr2: "test2",
		testDuration1: 1000 * time.Millisecond,
		testDuration2: 120000 * time.Millisecond,
		testInt1: 1,
		testInt2: 2,
		testBool1: true,
		testBool2: false,
	}
}

func (c *DynamicConfig) TestStr1(ctx context.Context) string {
	if c.client == nil {
		return c.testStr1
	}

resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(strKeyTestStr1).Cache(),
		time.Minute,
	).ToString()

	if err != nil {
		return c.testStr1
	}

	c.testStr1 = resp
	return c.testStr1
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
	if c.client == nil {
		return c.testStr2
	}

resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(strKeyTestStr2).Cache(),
		time.Minute,
	).ToString()

	if err != nil {
		return c.testStr2
	}

	c.testStr2 = resp
	return c.testStr2
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
	if c.client == nil {
		return c.testDuration1
	}

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
	if c.client == nil {
		return c.testDuration2
	}

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
	if c.client == nil {
		return c.testInt1
	}

resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(intKeyTestInt1).Cache(),
		time.Minute,
	).ToString()

	if err != nil {
		return c.testInt1
	}

	respInt, err := strconv.Atoi(resp)
	if err != nil {
		return c.testInt1
	}

	c.testInt1 = respInt
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
	if c.client == nil {
		return c.testInt2
	}

resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(intKeyTestInt2).Cache(),
		time.Minute,
	).ToString()

	if err != nil {
		return c.testInt2
	}

	respInt, err := strconv.Atoi(resp)
	if err != nil {
		return c.testInt2
	}

	c.testInt2 = respInt
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
	if c.client == nil {
		return c.testBool1
	}

resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(boolKeyTestBool1).Cache(),
		time.Minute,
	).ToString()

	if err != nil {
		return c.testBool1
	}

	c.testBool1 = resp == "true"
	return c.testBool1
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
	if c.client == nil {
		return c.testBool2
	}

resp, err := c.client.DoCache(
		ctx,
		c.client.B().Get().Key(boolKeyTestBool2).Cache(),
		time.Minute,
	).ToString()

	if err != nil {
		return c.testBool2
	}

	c.testBool2 = resp == "true"
	return c.testBool2
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

func (c *DynamicConfig) Config(ctx context.Context) ResponseDynamicConfig {
	return ResponseDynamicConfig{
		TestStr1: c.TestStr1(ctx),
		TestStr2: c.TestStr2(ctx),
		TestDuration1: c.TestDuration1(ctx),
		TestDuration2: c.TestDuration2(ctx),
		TestInt2: c.TestInt2(ctx),
		TestInt1: c.TestInt1(ctx),
		TestBool1: c.TestBool1(ctx),
		TestBool2: c.TestBool2(ctx),
	}
}

func (c *DynamicConfig) Update(ctx context.Context, req *RequestDynamicConfigUpdate) error {
	if req.TestStr1 != nil {
		if err := c.StoreTestStr1(ctx, *req.TestStr1); err != nil {
			return err
		}
	}
	if req.TestStr2 != nil {
		if err := c.StoreTestStr2(ctx, *req.TestStr2); err != nil {
			return err
		}
	}
	if req.TestDuration2 != nil {
		if err := c.StoreTestDuration2(ctx, *req.TestDuration2); err != nil {
			return err
		}
	}
	if req.TestDuration1 != nil {
		if err := c.StoreTestDuration1(ctx, *req.TestDuration1); err != nil {
			return err
		}
	}
	if req.TestInt1 != nil {
		if err := c.StoreTestInt1(ctx, *req.TestInt1); err != nil {
			return err
		}
	}
	if req.TestInt2 != nil {
		if err := c.StoreTestInt2(ctx, *req.TestInt2); err != nil {
			return err
		}
	}
	if req.TestBool1 != nil {
		if err := c.StoreTestBool1(ctx, *req.TestBool1); err != nil {
			return err
		}
	}
	if req.TestBool2 != nil {
		if err := c.StoreTestBool2(ctx, *req.TestBool2); err != nil {
			return err
		}
	}
	return nil
}
